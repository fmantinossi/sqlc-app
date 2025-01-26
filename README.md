# SQLC-APP Study Project

This Go application was created as a study project to explore the capabilities of [SQLC](https://sqlc.dev/) for database query generation and management. The project demonstrates the use of migrations for database creation, as well as transaction management in Go.

## Features

- **Database Migrations**: The `sql/migrations` folder contains scripts for initializing the database structure in MySQL.
  - `000001_init.up.sql`: Script for creating the initial database schema.
  - `000001_init.down.sql`: Script for rolling back the initial schema.

- **SQLC Integration**: The application uses SQLC to generate type-safe Go code for interacting with the database based on SQL queries defined in the `sql/queries` folder.
  - `query.sql`: Contains the SQL queries used for code generation.

- **Transaction Management**: The `cmd/runSQLTX` directory contains examples of using transactions to ensure data consistency.

## Project Structure

```
SQLC-APP
├── cmd
│   ├── runSQLC
│   │   └── main.go          # Example without transactions
│   └── runSQLTX
│       └── main.go          # Example with transactions
├── internal
│   └── db
│       ├── db.go            # Database connection setup
│       ├── models.go        # Generated models
│       └── query.sql.go     # Generated query methods
├── sql
│   ├── migrations
│   │   ├── 000001_init.up.sql  # Migration to set up database
│   │   └── 000001_init.down.sql # Rollback migration
│   └── queries
│       └── query.sql        # SQL queries for SQLC
├── .gitignore               # Git ignore file
├── docker-compose.yaml      # Docker setup for MySQL
├── go.mod                   # Go module file
├── go.sum                   # Dependencies checksum
├── Makefile                 # Automation tasks
├── README.md                # Project documentation
└── sqlc.yaml                # SQLC configuration
```

## Getting Started

1. **Clone the Repository**:
   ```bash
   git clone <repository_url>
   cd SQLC-APP
   ```

2. **Set Up the Database**:
   - Ensure you have MySQL running. You can use the `docker-compose.yaml` file to set up a local MySQL instance.
     ```bash
     docker-compose up -d
     ```
   - Run the migrations to set up the database schema.

3. **Generate SQLC Code**:
   - Use SQLC to generate Go code from the queries.
     ```bash
     sqlc generate
     ```

4. **Run the Application**:
   - Without transactions:
     ```bash
     go run cmd/runSQLC/main.go
     ```
   - With transactions:
     ```bash
     go run cmd/runSQLTX/main.go
     ```

## Prerequisites

- Go 1.20+
- SQLC installed ([Installation Guide](https://docs.sqlc.dev/en/stable/overview/install.html))
- MySQL or Docker (for running MySQL)

## Notes

- The project uses a `Makefile` to streamline tasks such as running migrations and generating code.
- The `internal/db` package handles the database connection and query execution.
- Transactional logic is demonstrated in the `cmd/runSQLTX` folder for scenarios requiring atomic operations.

## License

This project is for educational purposes only and is not intended for production use.