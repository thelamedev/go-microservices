package main

import (
	"errors"
	"log"

	"github.com/thelamedev/go-microservices-auth/utils"
)

type User struct {
	Id       string
	Name     string
	Username string
	Password string
	Email    string
}

type Database struct {
	Users    []User
	Sessions map[string]string
}

var Db *Database

func NewDatabase() *Database {
	return &Database{
		Users:    []User{},
		Sessions: map[string]string{},
	}
}

func SeedDatabase() {
	Db.Users = append(Db.Users, User{
		Id:       utils.NewId(),
		Name:     "micro-admin",
		Username: "admin",
		Password: "admin",
		Email:    "admin@gomicro.wtf",
	})

	log.Printf("Database seeded with %d users", len(Db.Users))
}

func FindUserByUsername(username string) (*User, error) {
	for _, item := range Db.Users {
		if item.Username == username {
			return &item, nil
		}
	}
	return nil, errors.New("user not found")
}

func FindUserById(id string) (*User, error) {
	for _, item := range Db.Users {
		if item.Id == id {
			return &item, nil
		}
	}
	return nil, errors.New("user not found")
}

func FindSessionByKey(sessionKey string) (string, error) {
	uid, ok := Db.Sessions[sessionKey]
	if !ok {
		return "", errors.New("session not found")
	}
	return uid, nil
}
