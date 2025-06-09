build:
	@go build -o bin/golang-backend-todo-api cmd/api/main.go

test:
	@go test -v ./...

run: build
	@./bin/golang-backend-todo-api

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down -v

docker-start:
	@docker-compose start

docker-stop:
	@docker-compose stop