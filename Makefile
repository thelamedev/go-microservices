.PHONY: gateway

gateway:
	go build -o ./bin/gateway ./services/gateway
	./bin/gateway

auth:
	go build -o ./bin/auth ./services/auth
	./bin/auth
