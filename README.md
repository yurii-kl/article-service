# Article Service

RESTful API service for managing articles.

## Prerequisites

- Go 1.24.3+
- PostgreSQL 15+
- Redis 7+

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Configure environment:
```bash
cp pkg/config/config.example.env .env
# Edit .env with your settings
```

3. Start services:
```bash
make docker-up
make db-migrate
```

4. Run service:
```bash
make run
```

## API

Swagger docs: `http://localhost:8080/api/swagger/index.html`

Endpoints:
- `POST /api/v1/article` - Create article
- `GET /api/v1/article/:id` - Get article
- `GET /api/private/health` - Health check
- `GET /api/private/metrics` - Prometheus metrics (view: `curl http://localhost:8080/api/private/metrics`)

## Testing

```bash
make test
make test-coverage
```

## Commands

- `make build` - Build binary
- `make run` - Run service
- `make test` - Run tests
- `make docker-up` - Start Docker services
- `make docker-down` - Stop Docker services
- `make db-migrate` - Run migrations
