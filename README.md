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
```

## 2. Set up PostgreSQL

without DOCKER

go run main.go

Using Docker:

docker run --name task-db -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_USER=postgres -e POSTGRES_DB=tasks_db \
  -p 5432:5432 -d postgres


## API DOCUMENTATION

1. Create Task
   POST /tasks
Content-Type: application/json

{
  "title": "Write unit tests",
  "status": "Pending"
}

2. Get Tasks with Pagination & Filter
   GET /tasks?status=Pending&page=1&limit=5

3. Update Task
   PUT /tasks/1
Content-Type: application/json

{
  "title": "Write more unit tests",
  "status": "Completed"
}



## Microservices Concepts

✅ Separation of Concerns
	•	Handler layer: only handles HTTP
	•	Service layer: business logic
	•	Repository layer: DB queries

✅ Single Responsibility Principle

Each package (handlers, services, repositories) focuses on a specific job.

✅ Statelessness

The microservice does not maintain user sessions — it’s stateless and horizontally scalable.

✅ Deployable as Unit

The service runs independently via Docker, making it container-ready and orchestration-friendly (Kubernetes, etc.).

✅ Extensibility

Supports easy addition of features like:
	•	Authentication middleware
	•	Logging/tracing
	•	gRPC communication between services

⸻

🔐 Future Improvements
	•	Add authentication (JWT/OAuth2)
	•	Use gRPC for inter-service communication
	•	Add unit and integration tests
	•	CI/CD with GitHub Actions or Docker Hub



  


