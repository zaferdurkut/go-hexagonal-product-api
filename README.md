# Go Hexagonal Product API

A simple Product API demonstrating Hexagonal Architecture (Ports & Adapters) in Go. It exposes HTTP endpoints using Gin and stores data in an in-memory repository.

## Prerequisites

- Go 1.22+ installed

## Getting Started

### Install dependencies

```bash
cd /go-hexagonal-product-api
go mod download
```

### Run the API

```bash
go run ./cmd/api
```

The server will start on port 8080.

## API

Base URL: `http://localhost:8080/api`

### Create product

- Method: POST
- Path: `/products`
- Body:
```json
{
  "name": "Apple"
}
```
- Responses:
  - 201 Created with product
  - 400 Bad Request on invalid body
  - 500 Internal Server Error on other failures

### List all products

- Method: GET
- Path: `/products`
- Responses:
  - 200 OK with array of products

### Get product by ID

- Method: GET
- Path: `/products/{id}`
- Responses:
  - 200 OK with product
  - 404 Not Found when product does not exist

### Get product details (concurrent example)

- Method: GET
- Path: `/products/{id}/details`
- Responses:
  - 200 OK with product detail:
```json
{
  "id": "<string>",
  "name": "<string>",
  "stock": 0,
  "fetchTime": "12ms"
}
```
  - 404 Not Found when product does not exist

## Project Structure

- `cmd/api/main.go`: Application entrypoint and HTTP server wiring
- `internal/core/domain`: Domain entities and errors
- `internal/core/ports`: Ports (interfaces) for repository and service
- `internal/core/services`: Application service implementing business logic
- `internal/adapters/http`: HTTP handlers (Gin)
- `internal/adapters/storage`: In-memory storage implementing repository port

## Notes

- Storage is in-memory and will reset on each run.
- Product IDs are generated as UUID strings on create.
