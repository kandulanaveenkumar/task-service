# Task Management Microservice (Go)

This is a simple Task Management System built in Go, designed as a microservice. It supports creating, reading, updating, and deleting tasks, along with pagination and status-based filtering. It includes Docker support and Swagger-based API documentation.

---

## Problem Breakdown

Build a microservice that:
- Supports CRUD operations on tasks
- Allows **pagination** and **filtering by task status**
- Follows **microservices principles** (single responsibility, separation of concerns)
- Can be deployed using **Docker**
- Exposes self-documenting **API using Swagger**

---

## Design Decisions

| Layer         | Responsibility                                         |
|---------------|--------------------------------------------------------|
| **Handler**   | HTTP route handling and JSON marshaling                |
| **Service**   | Business logic like validation and orchestration       |
| **Repository**| Database operations using GORM                         |
| **Model**     | Structs representing domain data (e.g., `Task`)        |
| **Router**    | Gorilla Mux-based route management                     |
| **Database**  | PostgreSQL via Docker or local                         |

- Uses **Gorilla Mux** for routing.
- Uses **GORM** for database abstraction.
- Uses **Swaggo** for Swagger documentation.
- Designed to be cloud/container ready with **Dockerfile**.

---

## Run the Service

### ▶️ Prerequisites
- [Go 1.21+](https://go.dev/)
- [Docker](https://www.docker.com/)
- PostgreSQL (local or Docker)

### 📦 1. Clone and Setup

```bash
git clone https://github.com/kandulanaveenkumar/task-service.git
cd task-service
go mod tidy
