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

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

  ```bash
  brew install golang-migrate
  ```

- [DB Docs](https://dbdocs.io/docs)

  ```bash
  npm install -g dbdocs
  dbdocs login
  ```

- [DBML CLI](https://www.dbml.org/cli/#installation)

  ```bash
  npm install -g @dbml/cli
  dbml2sql --version
  ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

  ```bash
  brew install sqlc
  ```

- [Gomock](https://github.com/golang/mock)

  ```bash
  go install github.com/golang/mock/mockgen@v1.6.0
  ```

### Setup infrastructure

- Create the bank-network

  ```bash
  make network
  ```

- Start postgres container:

  ```bash
  make postgres
  ```

- Create simple_bank database:

  ```bash
  make createdb
  ```

- Run db migration up all versions:

  ```bash
  make migrateup
  ```

- Run db migration up 1 version:

  ```bash
  make migrateup1
  ```

- Run db migration down all versions:

  ```bash
  make migratedown
  ```

- Run db migration down 1 version:

  ```bash
  make migratedown1
  ```

### Documentation

- Generate DB documentation:

  ```bash
  make db_docs
  ```

- Access the DB documentation at [this address](https://dbdocs.io/techschool.guru/simple_bank). Password: `secret`

### How to generate code

- Generate schema SQL file with DBML:

  ```bash
  make db_schema
  ```

- Generate SQL CRUD with sqlc:

  ```bash
  make sqlc
  ```

- Generate DB mock with gomock:

  ```bash
  make mock
  ```

- Create a new db migration:

  ```bash
  make new_migration name=<migration_name>
  ```

### How to run

- Run server:

  ```bash
  make server
  ```

- Run test:

  ```bash
  make test
  ```

## Deploy to kubernetes cluster

- [Install nginx ingress controller](https://kubernetes.github.io/ingress-nginx/deploy/#aws):

  ```bash
  kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml
  ```

- [Install cert-manager](https://cert-manager.io/docs/installation/kubernetes/):

  ```bash
  kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
  ```
