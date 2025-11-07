# South Park Messenger

A distributed system built with **Go**, **RabbitMQ**, and **Python**, following a hexagonal architecure approach.  
The project demonstrates message publishing, queueing, and consuming in a fun way — featuring quotes from South Park characters 🎭.

---

## 📘 Project Overview

This system is composed of three services:

| Service | Description | Port |
|----------|--------------|------|
| 🟣 **Go API Service** | Exposes a REST endpoint `/messages` that accepts JSON and publishes messages to a RabbitMQ queue (`southpark_messages`). Also serves the modern web UI. | 8080 |
| 🟢 **RabbitMQ Broker** | Message broker that stores messages until consumed. Includes a management UI. | 5672 (AMQP), 15672 (Management UI) |
| 🐍 **Python Consumer** | Listens to the same queue and prints received messages to the console. | – |

---

## 🐳 Running the Project

1. Build and start the stack (RabbitMQ, Go API + web UI, Python consumer):
   ```bash
   docker compose up --build
   ```
2. Wait until the logs show `go-api listening on :8080`. RabbitMQ and the consumer need a few extra seconds on the first run.
3. Confirm the API is healthy:
   ```bash
   curl http://localhost:8080/health
   ```
4. Open the web UI in a browser: http://localhost:8080  
   - Pick a character, type a quote, then hit **Send** (or turn on **Auto** to stream random quotes).  
   - Each message hits the Go API, which publishes to RabbitMQ, and the Python consumer prints the payload to its logs.
5. Watch processed messages (optional):
   ```bash
   docker compose logs -f python-consumer
   ```
### Stopping / cleaning up
- Stop containers: press `Ctrl+C` in the compose session or run `docker compose down`.
- Remove containers/images/network when done:
  ```bash
  docker compose down --volumes --remove-orphans
  ```