.PHONY: migrate-up migrate-down migrate-create migrate-version migrate-force help

DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_NAME ?= moviedb
DB_PASSWORD ?= haikus51w

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

help:
	@echo "Movie Database - Migration Commands"
	@echo "==================================="
	@echo "make migrate-up       Apply all pending migrations"
	@echo "make migrate-down     Rollback last migration"
	@echo "make migrate-create name=xxx   Create new migration"
	@echo "make migrate-version  Show current migration version"
	@echo "make migrate-force version=N   Force version"

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-version:
	migrate -path db/migrations -database "$(DB_URL)" version

migrate-force:
	migrate -path db/migrations -database "$(DB_URL)" force $(version)