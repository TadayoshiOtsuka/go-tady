SHELL=/bin/bash

.PHONY: $(shell egrep -oh ^[a-zA-Z0-9][a-zA-Z0-9_-]+: $(MAKEFILE_LIST) | sed 's/://')

build:
	docker compose -f docker-compose.test.yaml --env-file ./.env.test build

test:
	docker compose -f docker-compose.test.yaml --env-file .env.test up -d
	docker compose -f docker-compose.test.yaml --env-file .env.test exec server_test go test -v -coverpkg=./... ./.../test
