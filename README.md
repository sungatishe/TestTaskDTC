# TestTask

Application for managing orders and products developed in Go. It provides a REST API to manage orders and products, including error handling, logging, and caching to enhance performance.

## Project Structure

```plaintext
├── Dockerfile                  # Docker configuration for containerization
├── Makefile                    # Makefile for managing tasks
├── cmd/order/main.go           # Main application file
├── config                       # Application configuration
├── db                           # Database management and migrations
├── docker-compose.yaml         # Docker Compose for managing services
├── docs                         # Documentation and Swagger specification
├── go.mod                       # Go dependency management
├── go.sum                       # Go dependencies checksum
├── internal                     # Core logic of the application
├── pkg/utils                    # Utility functions like JWT handling
└── test                         # Unit tests for repositories and services
```

## Sections
- `cmd`: The main executable file to run the service.
- `config`: Stores application configurations, including the `config.yaml` file.
- `db`: Migrations for creating and altering the database schema and handling database configuration.
- `docs`: Swagger specification and documentation.
- `go.mod`: Go dependency management file.
- `go.sum`: Go dependencies checksum file.
- `internal`: Main application logic, including models, handlers, repositories, services, and routes.
- `pkg/utils`: Utility functions used across the project, such as JWT handling.
- `test`: Unit tests for services and repositories.

## Prerequisites

Before running the project, make sure you have the following software installed:

1. **Golang (>= 1.18)**
   - This project is built using Go, so you'll need Go installed to work with the code locally, run tests, or build the application from source.
   - You can install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

2. **Docker & Docker Compose**
   - Docker is used for containerization of the database (`postgres`) and the order service (`order-service`).
   - Docker Compose is used to orchestrate the multi-container setup.
   - You can install Docker and Docker Compose from: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/).

3. **golang-migrate (for database migrations)**
   - `golang-migrate` is used for managing database schema migrations.
   - You can install `golang-migrate` using the following command (if it's not available):

     ```bash
     brew install golang-migrate
     ```

     Or follow the official installation guide: [https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate).

4. **Swagger (Optional, if you want to regenerate documentation)**
   - Swagger is used for API documentation. It is required to run `make swag_init` to regenerate the API docs.
   - You can install Swagger using:

     ```bash
     brew install swag
     ```

     Or follow the installation instructions: [https://github.com/swaggo/swag](https://github.com/swaggo/swag).

---

### Additional Tools (Optional)

- **Make**
   - While not strictly required, **Make** is used to simplify various commands (e.g., running tests, setting up the database, building the application).
   - Make is typically pre-installed on most UNIX-like systems. If it's missing, install it via:

      - On Ubuntu: `sudo apt install make`
      - On macOS: `brew install make`

---

Once you have all these prerequisites installed, follow the instructions in the next section to set up and run the project.



## Running the Application

To set up and run the application, you can use Docker Compose and Makefile commands for convenience.

### Step-by-Step Guide

1. **Clone the repository:**

   Clone the repository to your local machine:
    ```bash
    git clone https://github.com/sungatishe/TestTaskDTC.git
    cd TestTask
    ```

2. **Set Up Environment Variables:**

   Make sure the `.env` file exists in the root of your project. This file contains database parameters and JWT secret. Example:

    ```dotenv
    APP_PORT=8080

    DB_HOST= << PUT YOUR DB HOST >>
    DB_PORT=5432
    DB_USER= << PUT YOUR DB USERNAME >>
    DB_PASSWORD= << PUT YOUR DB PASSWORD >>
    DB_NAME= # << PUT YOUR DB NAME >>
    DB_SSL_MODE=disable

    JWT_SECRET= << PUT YOUR JWT SECRETY KEY >>
    ```

3. **Build and Start the Services:**

   Use the Makefile to build and start the application services. These include the database and the main order service.

   To build everything and start the containers, run:
    ```bash
    make build
    ```

   Alternatively, you can run the services without rebuilding with:
    ```bash
    make up
    ```

4. **Database Setup (if necessary):**

   If the database has not been set up yet, run the following command to create the database:
    ```bash
    make createdb
    ```

   If you need to apply database migrations to initialize the schema or update the database, run:
    ```bash
    make migrateup
    ```


5. **Stopping the Application:**

   To stop the running services, you can use the following command:
    ```bash
    make down
    ```

   This will stop and remove the containers.

---

### Quick Commands

Below are the most common Makefile commands:

```bash
make build            # Build and start the services
make up               # Start the services (without rebuild)
make down             # Stop the running services
make createdb         # Create the PostgreSQL database
make dropdb           # Drop the PostgreSQL database
make migrateup        # Run database migrations
make migratedown      # Roll back database migrations
make swag_init        # Generate Swagger documentation
make test             # Run the unit tests
```
## Docker Compose Overview
The application uses Docker Compose to manage the services. These include:

#### PostgreSQL Database (db):

- A PostgreSQL 15 container that is configured using environment variables from the .env file.
- Ports: 5432 (database port)

### Order Service (order-service):

- The Go-based service running the order management API.
- Ports: 8080 (API access)
- Depends on the db service, ensuring the database is healthy before starting.

## API Documentation
The API documentation is generated using Swagger and is available at `/swagger/index.html`. You can visit this URL in the browser while the app is running.

## Tech Stack

- **Go**: The main programming language.
- **PostgreSQL**: Relational database for data storage.
- **Docker**: For containerization of the project.
- **Swagger**: For generating API documentation.
