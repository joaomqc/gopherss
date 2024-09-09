SERVER_DIR = server

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	cd $(SERVER_DIR) &&	swag fmt
	cd $(SERVER_DIR) && swag init

server: swagger
	cd $(SERVER_DIR) && go run .