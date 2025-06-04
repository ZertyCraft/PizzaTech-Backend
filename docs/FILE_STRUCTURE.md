# Project File Structure

Below is the high-level overview of the directory layout and purpose of each folder/file:

````

.
├── cmd
│   └── server
│       └── main.go              # Entry point for the Gin server
├── config
│   └── config.go                # Loads environment variables into a Config struct
├── internal
│   ├── delivery
│   │   └── http
│   │       ├── handlers         # HTTP handler implementations
│   │       │   ├── auth\_handler.go
│   │       │   ├── pizza\_handler.go
│   │       │   ├── order\_handler.go
│   │       │   ├── profile\_handler.go
│   │       │   └── stats\_handler.go
│   │       └── middlewares      # Gin middleware definitions
│   │           ├── auth\_middleware.go
│   │           └── logger\_middleware.go
│   ├── domain
│   │   ├── models               # GORM models
│   │   │   ├── user.go
│   │   │   ├── pizza.go
│   │   │   └── order.go
│   │   └── repositories         # Repository interfaces
│   │       ├── user\_repository.go
│   │       ├── pizza\_repository.go
│   │       └── order\_repository.go
│   ├── infrastructure
│   │   ├── logger               # Logrus logger initialization
│   │   │   └── logger.go
│   │   └── persistence          # GORM repository implementations
│   │       ├── db.go
│   │       ├── user\_repository.go
│   │       ├── pizza\_repository.go
│   │       └── order\_repository.go
│   ├── services                 # Business logic / service layer
│   │   ├── auth\_service.go
│   │   ├── pizza\_service.go
│   │   ├── order\_service.go
│   │   └── statistics\_service.go
│   └── di                       # Dependency injection with Uber Dig
│       └── di.go
├── .env.example                 # Example environment variables
├── docker-compose.yml           # Docker Compose definition for app + DB
├── go.mod                       # Go modules file
└── go.sum                       # Go dependency checksums

````

## Descriptions

- **cmd/server/**: The main application entry. Starts the Gin engine via DI.
- **config/**: Centralized config loader that reads `.env`.
- **internal/delivery/http/handlers/**: Contains Gin handlers for each route. They translate HTTP requests to service calls and format responses.
- **internal/delivery/http/middlewares/**: Custom Gin middlewares (logging, authentication).
- **internal/domain/models/**: GORM models for User, Pizza, Order, and OrderItem.
- **internal/domain/repositories/**: Interfaces that define repository operations to decouple the service layer from persistence.
- **internal/infrastructure/logger/**: Initializes and configures a Logrus logger.
- **internal/infrastructure/persistence/**: GORM-based implementations of each repository interface. Also, `db.go` initializes the PostgreSQL connection.
- **internal/services/**: Business logic—validates inputs, manages transactions, issues repository calls, and returns results to handlers.
- **internal/di/**: Wires all dependencies (config, logger, DB, repositories, services, and handlers) using Uber Dig.