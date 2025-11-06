# SSP2 South Park – Microservices (Go + RabbitMQ + Python)

## Services
- **go-api**: HTTP API (Hexagonal Architecture) – will publish to RabbitMQ
- **python-consumer**: consumes messages from RabbitMQ and logs them
- **rabbitmq**: message broker (ports: 5672, 15672 UI)

## Quick start

docker compose up --build
# health check
curl http://localhost:8080/health
