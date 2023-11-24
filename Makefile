#!make
include .env

all:

.PHONY: migrate-up
migrate-up:
	migrate -database $(DB_DSN) -path migrations up

.PHONY: migrate-down
migrate-down:
	migrate -database $(DB_DSN) -path migrations down