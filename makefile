# up the docker containers
up:
	docker-compose -f deployment/docker-compose.yml up -d

# down the docker containers
down:
	docker-compose -f deployment/docker-compose.yml down

# run the application
run:
	go run main.go

lint:
	golangci-lint run ./...

.PHONY: up down run lint