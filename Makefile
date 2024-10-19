.PHONY: gateway

gateway:
	go build -o ./bin/gateway ./services/gateway
	./bin/gateway
