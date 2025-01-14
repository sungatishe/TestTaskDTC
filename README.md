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

5. **Health Check:**

   The `order-service` will automatically wait for the `db` to be ready due to `depends_on` configuration in `docker-compose.yaml`.

   If you'd like to manually check if your database is reachable from within the `order-service` container, run:
    ```bash
    docker exec -it order-service pg_isready -h order_db -U ${DB_USER} -d ${DB_NAME}
    ```

6. **Testing the Application (Optional):**

   To run the unit tests for repositories and services, use the following command:
    ```bash
    make test
    ```

7. **Stopping the Application:**

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
