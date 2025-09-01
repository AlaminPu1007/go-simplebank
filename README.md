# Simple Bank

A backend API service for a simple banking system built with **Go**, **PostgreSQL**, and **Gin**.  
It supports user management, account operations, and money transfers with secure authentication.

```

## 🚀 Features

-   User registration & login with hashed passwords
-   JWT authentication & authorization
-   CRUD operations for bank accounts
-   Transfer money between accounts
-   Database migrations with `golang-migrate`
-   Unit & integration tests
-   Docker & Docker Compose support
-   Makefile for automation

```

## 🛠️ Tech Stack

-   [Go](https://golang.org/) – Programming language
-   [Gin](https://github.com/gin-gonic/gin) – Web framework
-   [PostgreSQL](https://www.postgresql.org/) – Relational database
-   [golang-migrate](https://github.com/golang-migrate/migrate) – Database migrations
-   [Docker](https://www.docker.com/) – Containerization
-   [Testify](https://github.com/stretchr/testify) – Testing framework

```

## 📂 Project Structure

.
├── api/ # HTTP API endpoints
├── db/ # Database SQL, queries & migrations
│ ├── migration/ # Migration files
│ ├── sqlc/ # Auto-generated query code (sqlc)
│ └── schema.sql # DB schema
├── util/ # Utility functions (e.g., password hashing, random strings)
├── token/ # JWT & token management
├── main.go # Application entry point
├── Makefile # Common tasks automation
└── go.mod / go.sum # Go dependencies


```
