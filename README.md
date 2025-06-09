# 📝 Todo API

A RESTful API for a user-based todo list with JWT authentication, built using
**Go**, **PostgreSQL**, and **Docker**.

---

## 🚀 Features

- User registration and authentication (JWT-based)
- Create, read, update, and delete todos
- PostgreSQL integration via Docker
- Environment-based configuration
- Clean project structure

---

## 🧰 Prerequisites

Make sure you have the following installed:

- [Go](https://go.dev/doc/install) (v1.24.4 or later)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 📦 Setup & Run

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/golang-todo-api.git
cd golang-todo-api
```

### 2. Create a .env File

```bash
DB_HOST=localhost
DB_PORT=<port>
DB_USER=<username>
DB_PASSWORD=<user_password>
DB_NAME=<database_name>
SERVER_PORT=<server_port>
```

### 3. PostgreSQL via Docker

#### To create and run postgres docker container:

```bash
make docker-up
```

#### To stop and remove the containers and volumes:

```bash
make docker-down
```

#### To start/stop containers without removing volumes:

```bash
make docker-start
make docker-stop
```

### 4. Build & Run the API

```bash
make run
```

### 5. To run test (optional)

```bash
make test
```

## Project folder structure

```bash
todo-api/
├── cmd/
│   └── api/
│       └── main.go              # Entry point for the application
├── internal/
│   ├── config/                  # Configuration management
│   │   └── config.go
│   ├── handlers/                # HTTP handlers for API endpoints
│   │   ├──
│   ├── models/                  # Data models (structs) for database entities
│   │   ├──
│   ├── repository/              # Database operations (CRUD)
│   │   ├──
│   ├── services/                # Business logic
│   │   ├──
│   └── middleware/              # Middleware for authentication, logging, etc.
│       └──
├── pkg/                         # Shared utilities and packages
│   └── database/
│       └── database.go          # Database connection setup
├── docker-compose.yml           # Docker configuration for PostgreSQL
├── .env                         # Environment variables
├── go.mod                       # Go module file
├── go.sum                       # Go module dependencies
├── Makefile                     # To run scripts or automate tasks
└── README.md                    # Project documentation
```
