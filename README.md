# Chat Backend Service 🚀

A real-time chat backend built in Go with Kafka and WebSocket integration.

---

## 🔍 Overview

This service handles:

- Producing and consuming chat messages using **Kafka**
- Maintaining multiple consumer groups (e.g., `chat-service`, `websocket-service`)
- Delivering messages to clients via **WebSocket**

---

## 📦 Features

- **Producer**: Emits user messages to Kafka topics (e.g., `chat.messages`)
- **Consumer**: Reads from Kafka, writes to database, processes receipts
- **WebSocket Bridge**: Relays real‑time messages to connected clients
- Flexible topic structure to support:
    - `chat.messages`
    - `chat.read`
    - `chat.typing`
    - `chat.status`

---

## 🛠️ Getting Started

### Prerequisites

- Go (v1.20+)
- Kafka cluster (local or hosted)
- Redis (for WebSocket session management)
- (Optional) PostgreSQL / MongoDB for message storage

### Configuration

Create a `.env` file specifying:

```dotenv
KAFKA_BROKERS=localhost:9092
CHAT_SERVICE_GROUP=chat-service-group
WS_SERVICE_GROUP=websocket-service-group
DB_URL=...
REDIS_URL=...
