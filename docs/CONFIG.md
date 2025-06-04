# Configuration

The application uses environment variables to configure the database connection, JWT secret, and server port. These variables are loaded via `github.com/joho/godotenv` from a `.env` file in the project root.

## .env File

Copy `.env.example` to `.env` and update values as needed:


| Variable     | Description                                 | Example           |
| ------------ | ------------------------------------------- | ----------------- |
| `DB_HOST`    | Hostname or IP of PostgreSQL server         | `localhost`       |
| `DB_PORT`    | Port on which PostgreSQL listens            | `5432`            |
| `DB_USER`    | PostgreSQL username                         | `postgres`        |
| `DB_PASSWORD`| PostgreSQL password                         | `password`        |
| `DB_NAME`    | Name of the PostgreSQL database             | `pizzadb`         |
| `JWT_SECRET` | Secret key used to sign JWT tokens          | `supersecretkey`  |
| `SERVER_PORT`| Port for the Gin HTTP server                | `8080`            |


### Variables

- **DB_HOST**  
  Hostname or IP address of the PostgreSQL server.  
  - Example: `localhost`, `db` (in Docker Compose)

- **DB_PORT**  
  Port on which PostgreSQL listens (default: `5432`).  

- **DB_USER**  
  Username for connecting to PostgreSQL.  

- **DB_PASSWORD**  
  Password for the specified `DB_USER`.  

- **DB_NAME**  
  Name of the PostgreSQL database to use. The application will attempt to connect to this database and auto-migrate tables.

- **JWT_SECRET**  
  A secret key used to sign and verify JWT tokens. Keep this value secure and do not commit it to version control.

- **SERVER_PORT**  
  Port on which the Gin HTTP server will listen.  
  - Example: `8080`

## Loading Configuration

The `config.Load()` function (defined in `config/config.go`) reads the variables from `.env`, validates required fields, and returns a `*config.Config` struct:

```go
type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
    ServerPort string
}

func Load() (*Config, error) {
    godotenv.Load()
    cfg := &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        JWTSecret:  os.Getenv("JWT_SECRET"),
        ServerPort: os.Getenv("SERVER_PORT"),
    }
    if cfg.DBHost == "" || cfg.JWTSecret == "" || cfg.ServerPort == "" {
        return nil, fmt.Errorf("missing required env variables")
    }
    return cfg, nil
}
````

## Customizing Configuration

* To change the database host (e.g., when running via Docker Compose), update `DB_HOST=db`.
* To switch ports or database credentials, simply modify the values in `.env` and restart the service.
* The server port can be any unused port; ensure it matches the port mapping in `docker-compose.yml` if using containers.
