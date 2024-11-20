package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/thelamedev/go-microservices-auth/utils"
)

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	// read the body as json
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	// validate the body
	if body.Username == "" || body.Password == "" {
		utils.WriteError(w, errors.New("username and password are required"), 500)
		return
	}

	// check the user in database for login
	user, err := FindUserByUsername(body.Username)
	if err != nil {
		utils.WriteError(w, errors.New("credentails do not match"), 500)
		return
	}
	if user.Password != body.Password {
		utils.WriteError(w, errors.New("credentails do not match"), 500)
		return
	}

	// create session for the user
	skey := utils.NewId()
	Db.Sessions[skey] = user.Id

	// return the session key as json
	utils.WriteJSON(w, map[string]any{
		"message":    fmt.Sprintf("Welcome back, %s!", user.Name),
		"sessionKey": skey,
	}, 200)
}

func HandleUserSignup(w http.ResponseWriter, r *http.Request) {
	// read the body as json
	var body struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	// validate the body
	if body.Username == "" || body.Password == "" || body.Name == "" || body.Email == "" {
		utils.WriteError(w, errors.New("missing required fields"), 500)
		return
	}

	// check the user in database for login
	user, err := FindUserByUsername(body.Username)
	if err == nil || user.Id != "" {
		utils.WriteError(w, errors.New("username is taken"), 500)
		return
	}

	newUser := User{
		Id:       utils.NewId(),
		Name:     body.Name,
		Email:    body.Email,
		Username: body.Username,
		Password: body.Password,
	}

	// create session for the user
	Db.Users = append(Db.Users, newUser)

	skey := utils.NewId()
	Db.Sessions[skey] = user.Id

	// return the session key as json
	utils.WriteJSON(w, map[string]any{
		"message":    fmt.Sprintf("Welcome, %s!", user.Name),
		"sessionKey": skey,
	}, 200)
}

func HandleUserLogout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.WriteError(w, errors.New("no session key provided"), 400)
		return
	}

	tokenParts := strings.SplitN(authHeader, " ", 2)
	if len(tokenParts) != 2 {
		utils.WriteError(w, errors.New("malformed authorization header"), 400)
		return
	}

	sessionKey := tokenParts[1]
	_, err := FindSessionByKey(sessionKey)
	if err != nil {
		utils.WriteError(w, err, 401)
		return
	}

	delete(Db.Sessions, sessionKey)

	utils.WriteJSON(w, map[string]any{
		"message": "Logout successful",
	}, 200)
}

func HandleUserVerify(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.WriteError(w, errors.New("no session key provided"), 400)
		return
	}

	tokenParts := strings.SplitN(authHeader, " ", 2)
	if len(tokenParts) != 2 {
		utils.WriteError(w, errors.New("malformed authorization header"), 400)
		return
	}

	sessionKey := tokenParts[1]
	userId, err := FindSessionByKey(sessionKey)
	if err != nil {
		utils.WriteError(w, err, 401)
		return
	}

	user, err := FindUserById(userId)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	utils.WriteJSON(w, map[string]any{
		"message": "Session Verified",
		"profile": user,
	}, 200)
}
