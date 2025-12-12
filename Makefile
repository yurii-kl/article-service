ENVFILE ?= .env

ifneq (,$(wildcard $(ENVFILE)))
    include $(ENVFILE)
    export
endif

BINARY_NAME=article-service
BUILD_DIR=build

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

.PHONY: build
build:
	@$(GOMOD) tidy
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api

.PHONY: run
run:
	@$(GOBUILD) -o $(BINARY_NAME) ./cmd/api
	@./$(BINARY_NAME)

.PHONY: test
test:
	@$(GOTEST) -v ./...

.PHONY: test-coverage
test-coverage:
	@$(GOTEST) -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -func=coverage.out

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: lint-fix
lint-fix:
	@golangci-lint run --fix

.PHONY: docker-build
docker-build:
	@docker build -t $(BINARY_NAME):latest .

.PHONY: docker-up
docker-up:
	@docker compose up -d

.PHONY: docker-down
docker-down:
	@docker compose down

.PHONY: db-migrate
db-migrate:
	@migrate -path pkg/postgres/migrations/ \
		-database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSL)" \
		-verbose up

.PHONY: db-rollback
db-rollback:
	@migrate -path pkg/postgres/migrations/ \
		-database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSL)" \
		-verbose down 1
