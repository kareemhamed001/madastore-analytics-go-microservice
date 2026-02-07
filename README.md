# AnalyticsService

AnalyticsService is a Go service that powers analytics dashboards and exposes both HTTP and gRPC APIs. It reads visit data from MySQL, uses Redis for caching, and provides metrics and health endpoints for production observability.

## What this service does

- Serves analytics dashboards (top products, visits per day/month/country/city, and Egypt-specific stats)
- Exposes both HTTP REST and gRPC APIs
- Caches hot analytics in Redis
- Refreshes caches in the background on a schedule
- Exposes Prometheus metrics

## Architecture overview (human-friendly)

- **cmd/main.go** boots the application and manages graceful shutdown
- **internal/app** wires configuration, DB, cache, HTTP routes, gRPC server, and workers
- **internal/services** contains business logic
- **internal/repository** talks to MySQL with timeouts
- **internal/middleware** adds request IDs, structured logs, and timeouts
- **internal/worker** runs background jobs (cache refresher)
- **proto/** and **common/genproto/** define gRPC types

## Requirements

- Go 1.25+
- MySQL
- Redis

## Configuration

The service reads environment variables (supports both legacy and new names).

Common variables:

- APP_NAME
- ENVIRONMENT (development|production)
- SERVER_PORT
- GRPC_PORT
- DATABASE_DSN
- API_KEY
- REDIS_ADDR
- REDIS_PASSWORD
- REDIS_DB
- REQUEST_TIMEOUT
- REPO_QUERY_TIMEOUT
- CACHE_REFRESH_INTERVAL
- SHUTDOWN_TIMEOUT

Example:

```
export ENVIRONMENT=development
export SERVER_PORT=8080
export GRPC_PORT=9090
export DATABASE_DSN="user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true"
export API_KEY="your-api-key"
export REDIS_ADDR=127.0.0.1:6379
```

## Run the service

```
go run ./cmd
```

## Migrations

A simple migration CLI is included.

```
export MIGRATE_DATABASE_URL="mysql://user:pass@tcp(127.0.0.1:3306)/dbname"

go run ./cmd/migrate -dir migrations -action up
```

Supported actions: up, down, version, force

## HTTP Endpoints

- GET /health
- GET /metrics
- GET /api/v1/analytics
- GET /api/v1/analytics/top-products-visits
- GET /api/v1/analytics/visits-per-day
- GET /api/v1/analytics/visits-per-month
- GET /api/v1/analytics/visits-per-country
- GET /api/v1/analytics/visits-per-city
- GET /api/v1/analytics/visits-per-city-today
- GET /api/v1/analytics/visits-from-egypt-per-day
- GET /api/v1/analytics/visits-from-other-countries-per-day
- GET /api/v1/analytics/visits-from-egypt-per-hour-past-month
- GET /api/v1/analytics/visits-from-egypt-per-hour-today

All /api/v1 routes require the X-API-KEY header.

## gRPC

The gRPC server runs on GRPC_PORT and exposes analytics service methods matching the REST endpoints.

## Observability

- Structured JSON logs with request IDs
- Prometheus metrics at /metrics
- Request timeouts and repository timeouts to keep latency bounded

## Notes

- Redis is optional; if unavailable, the service will still run without cache.
- Background cache refresh runs on CACHE_REFRESH_INTERVAL.
