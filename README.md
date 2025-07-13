# Chat Backend (WIP) 💬

A distributed chat backend system written in Go, following microservice architecture with Kafka as the message backbone and WebSocket support for real-time communication.

---

## 🚧 Project Status

> 🛠️ This project is currently under active development.

So far, the system includes:

- `chat_service`: Handles chat message production and storage.
- `user_service`: Manages users and user-related operations.
- `websocket_service`: Pushes real-time messages to clients using WebSockets.
- All services communicate **asynchronously** through **Kafka topics**.

---

## 📦 Architecture Overview

user/client]
⬇️ WebSocket
[websocket_service]
⬇️ Kafka (consumer)
[chat_messages topic]
⬆️ Kafka (producer)
[chat_service]
⬅️ REST/GRPC/API
[user_service]



### 🔗 Communication
- Services are **decoupled** and communicate via **Kafka topics**.
- `websocket_service` consumes Kafka messages and delivers them to connected users.
- `chat_service` both produces and consumes from Kafka.
- `user_service` currently handles authentication and user data.

---

## 🧩 Services

### `chat_service`
- Produces messages to Kafka
- Consumes for post-processing or storage
- Handles chat business logic

### `websocket_service`
- Maintains active WebSocket connections
- Subscribed to Kafka topics like `chat.messages`
- Delivers messages to clients in real-time

### `user_service`
- User creation, login, and auth (in progress)
- Will likely use REST or gRPC for communication

---

## 🧰 Tech Stack

- **Language**: Go
- **Messaging**: Kafka (via [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go))
- **Real-time**: WebSocket
- **Data Storage**: TBD (likely PostgreSQL, Redis)
- **Architecture**: Microservices

---

## 🔧 Running Locally

> Setup instructions coming soon.

For now:
```bash
go mod download
# Example
go run cmd/chat-service/main.go
go run cmd/websocket-service/main.go
```

# Project Structure (WIP)
```bash
chat_backend/
├── internal/
│   │──messages
│   ├── chat-service/
│   ├── websocket-service/
│   └── user-service/
├── cmd/
│   ├── server/main.go
└── README.md
