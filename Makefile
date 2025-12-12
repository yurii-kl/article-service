ENVFILE ?= .env

ifneq (,$(wildcard $(ENVFILE)))
    include $(ENVFILE)
    export
endif

.PHONY: build
build:
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.version=$(Version)" -o app

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run-new
run-new:
	go build
	./article-service

.PHONY: docker-build
docker-build:
	docker build -t article-service:latest .

.PHONY: docker-up
docker-up:
	docker compose up -d

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: docker-logs
docker-logs:
	docker compose logs -f

.PHONY: docker-restart
docker-restart: docker-down docker-up

.PHONY: db-migrate
db-migrate:
	migrate -path pkg/postgres/migrations/ \
	   -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSL)" \
	   -verbose up

.PHONY: db-rollback
db-rollback:
	migrate -path pkg/postgres/migrations/ \
	   -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=$(POSTGRES_SSL)" \
	   -verbose down 1


