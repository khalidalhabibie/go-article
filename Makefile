.PHONY: clean critic security lint test build run
	
include .env
export

# APP_NAME = ${APP_NAME_ENV}
APP_NAME=goarticle
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL =postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${DB_SSL_MODE}



# ==============================================================================
# checking  golang tools


critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

	# go test -short  ./...



# ==============================================================================
# Build and run 


build:
	@echo build ....
	sudo docker-compose build

run:
	@echo run ....
	sudo docker-compose up

stop:
	@echo stop ....
	sudo docker-compose stop

remove: 
	@echo remove ....
	sudo docker-compose rm
	



# ==============================================================================
# Migration

migrate.up:	
	@echo migration up ....
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	@echo migration down 1 ...
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down 1
	
migrate.force $(version):
	@echo migration force to clean databse version $(version)
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)


migrate.create $(file): 
	@echo migration create $(file)
	migrate create -ext sql -dir $(MIGRATIONS_FOLDER) -seq $(file)


# ==============================================================================
# Swagger

swagger:
	@echo Starting swagger generating
	swag init -g **/**/*.go


# ==============================================================================
# Docker-compose

docker.build:
	@echo build or rebuil docker compose
	sudo docker-compose up --build