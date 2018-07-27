.PHONY: build dev bin local-db local-env clean-local-env unit-test integration-test test clean-env-force

bin:
	go build -o bin/migrator ./cmd/migrator

build: bin
	go build -o server cmd/server/*.go

clean-env-force:
	docker ps -a | grep example-go | awk '{ print $1 }' | xargs docker kill > /dev/null

local-db: bin
	eval "docker-compose -f localdb-docker-compose.yaml down"
	eval "docker-compose -f localdb-docker-compose.yaml up -d"

local-env: local-db
	@cat .env_migrator.yaml.example > .env_migrator.yaml
	@cat .env.example > .env
	@echo "Waiting for database connection..."
	@while ! docker exec examplego_db_1 pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	bin/migrator up

clean-local-env:
	eval "docker-compose -f localdb-docker-compose.yaml down"

integration-test:
	go test ./... -tags=integration -count=1

unit-test:
	go test ./... -tags=unit -count=1

test: unit-test integration-test

dev: build
	ENV=local ./server; rm server
