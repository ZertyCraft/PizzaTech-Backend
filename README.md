# PizzaTech Backend

> **Academic context**  
> This project was developed as part of the Master's program **“Manager Opérationnel d'Activité”** at **Ensemble Scolaire Saint Louis** in **Crest (Drôme, France)**.  
> The repository is public and intended for educational and portfolio purposes.

PizzaTech is a backend service for managing a pizza restaurant. It provides RESTful endpoints for user authentication, pizza catalog management, order processing, and statistics. The service is built in Go using the Gin framework, GORM for PostgreSQL integration, Logrus for logging, and Dig (Uber’s DI) for dependency injection. Configuration is managed through a `.env` file.

## Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Environment Variables](#environment-variables)
  - [Running with Docker Compose](#running-with-docker-compose)
  - [Building from Source](#building-from-source)
- [Project Structure](#project-structure)
- [Routing](#routing)
- [Configuration](#configuration)
- [License](#license)

## Features

- **Authentication**
  - Register and login for Customers, Admins, and Workers.
  - JWT-based authorization with role enforcement.
- **Pizza Catalog**
  - Create, list, retrieve, update, and delete pizzas (authenticated).
- **Orders**
  - Customers can place orders and view their order history.
  - Workers/Admins can update order statuses.
  - Admins can retrieve statistics (total orders).
- **Profile**
  - Customers can view their order history.
- **Dependency Injection**
  - All services, repositories, and handlers are wired with Uber Dig.
- **Logging**
  - Request logging via Logrus with JSON formatter.
- **Persistence**
  - PostgreSQL database managed via GORM.

## Technology Stack

- **Language**: Go 1.24
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io)
- **Database Driver**: `gorm.io/driver/postgres`
- **Dependency Injection**: [Uber Dig](https://github.com/uber-go/dig)
- **Logging**: [Logrus](https://github.com/sirupsen/logrus)
- **Environment Variables**: [godotenv](https://github.com/joho/godotenv)
- **JWT**: [jwt-go](https://github.com/dgrijalva/jwt-go)
- **Containerization**: Docker & Docker Compose

## Getting Started

### Prerequisites

- Go 1.24 or higher installed
- Docker & Docker Compose

### Environment Variables

Create a `.env` file in the project root based on the following template:

```

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=pizzadb
JWT_SECRET=supersecretkey
SERVER_PORT=8080

```

- `DB_HOST` – Hostname or IP of the PostgreSQL server.
- `DB_PORT` – Port on which PostgreSQL listens (default: 5432).
- `DB_USER` – PostgreSQL username.
- `DB_PASSWORD` – PostgreSQL password.
- `DB_NAME` – Name of the database to use.
- `JWT_SECRET` – Secret key for signing JWT tokens.
- `SERVER_PORT` – Port on which the Gin server will run.

### Running with Docker Compose

1. Ensure Docker and Docker Compose are installed.
2. Copy `.env.example` to `.env` and update values if needed.
3. From the project root, run:

   ```bash
   docker-compose up --build
   ```

4. The backend will be accessible at `http://localhost:8080`.

### Building from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/ZertyCraft/PizzaTech-Backend
   cd PizzaTech-Backend
   ```

2. Create a `.env` file (see [Environment Variables](#environment-variables)).
3. Install dependencies:

   ```bash
   go mod download
   ```
4. Build the binary:

   ```bash
   go build -o pizzatech ./cmd/server
   ```
5. Run the server:

   ```bash
   ./pizzatech
   ```

## Project Structure

See [FILE_STRUCTURE.md](./docs/FILE_STRUCTURE.md) for a detailed breakdown of directories and files.

## Routing

See [ROUTING.md](./docs/ROUTING.md) for a complete list of endpoints, HTTP methods, path parameters, and required roles.

## Configuration

See [CONFIG.md](./docs/CONFIG.md) for details on environment variables, config loading, and customizing settings.
