# Booking System — Event-Driven Architecture with Go & Kafka

This project showcases a simple **event-driven booking system** built in **Go** using **Apache Kafka** as the messaging backbone.
It is structured into multiple independent microservices to demonstrate scalability, decoupling, and asynchronous communication.

---

## Architecture Overview

- **Booking Service**:  
  Creates bookings and emits `booking-created` events to Kafka.
  
- **Payment Service**:  
  Listens for `booking-created` events, processes payments, and emits `payment-confirmed` events.

- **Notification Service**:  
  Listens for `payment-confirmed` events and sends notifications to users.

- **Kafka Broker**:  
  Handles all the messaging between services.

- **Zookeeper**:  
  Coordination service required by Kafka.

---

## Project Structure

```
booking-system/
│
├── app/
│   ├── kafka/         # Common Kafka producer, consumer, and topic manager logic
│   └── models/        # Common domain models (BookingCreated, PaymentCompleted, BookingConfirmed)
│
├── booking-service/   # Handles booking creation
│   ├── service/       
│   ├── go.mod
│   └── go.sum
│
├── payment-service/   # Handles payment processing
│   ├── service/
│   ├── go.mod
│   └── go.sum
│
├── notification-service/ # Handles notifications
│   ├── service/
│   ├── go.mod
│   └── go.sum
│
├── docker-compose.yml # Development setup with Kafka, Zookeeper, and all services
└── README.md
```

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/rodolfodiazr/booking-system.git
cd booking-system
```

### 2. Build and run services with Docker Compose

```bash
docker-compose up --build
```

This will:
- Start Zookeeper and Kafka
- Build each microservice
- Launch services on:
  - Booking Service → [localhost:8080](http://localhost:8080)
  - Payment Service → [localhost:8081](http://localhost:8081)
  - Notification Service → [localhost:8082](http://localhost:8082)

Kafka Broker will be available at **localhost:9092**.

---

## Environment Variables

Each service requires:

| Variable | Description |
|:---------|:------------|
| `KAFKA_BROKER_ADDRESS` | Kafka broker address (default: `kafka:9092`) |

> These are already configured inside the `docker-compose.yml`.

---

## Tech Stack

- **Go 1.24.1**
- **Kafka**
- **Docker Compose** for local orchestration
- **Kafka-go** client library