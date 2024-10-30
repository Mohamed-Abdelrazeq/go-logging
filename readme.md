# Go Logging Service
This project is a logging service written in Go. It provides an API to create and retrieve log records stored in an SQLite database.

## Project Structure

## Getting Started

### Prerequisites

- Go 1.16 or later
- SQLite3

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/MohamedAbdelrazeq/go-logging.git
    cd go-logging
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Create the SQLite database:

    ```sh
    sqlite3 logger-db.sqlite < schema.sql
    ```

### Running the Application

To start the application, run:

```sh
go run main.go
```

The server will start on [http://localhost:8080](http://localhost:8080).

## API Endpoints

### Health Check

- **URL:** `/health`
- **Method:** `GET`
- **Description:** Check the health of the service.

### Create Log Record

- **URL:** `/create-record`
- **Method:** `POST`
- **Description:** Create a new log record.
- **Request Body:**

    ```json
    {
        "level": "INFO",
        "message": "This is a log message",
        "timestamp": "2023-10-01T12:00:00Z"
    }
    ```

### Get Log Records

- **URL:** `/get-records`
- **Method:** `GET`
- **Description:** Retrieve log records with pagination.
- **Query Parameters:**
  - `page` (optional): Page number (default is 0).

### Get Log Record by ID

- **URL:** `/get-record`
- **Method:** `GET`
- **Description:** Retrieve a log record by its ID.
- **Query Parameters:**
  - `id`: Log record ID.

### Get Log Records by Level

- **URL:** `/get-records-by-level`
- **Method:** `GET`
- **Description:** Retrieve log records by log level with pagination.
- **Query Parameters:**
  - `level`: Log level (e.g., INFO, ERROR).
  - `page` (optional): Page number (default is 0).

### Get Log Records by Date Range

- **URL:** `/get-records-by-date-range`
- **Method:** `GET`
- **Description:** Retrieve log records within a specific date range.
- **Query Parameters:**
  - `startDate`: Start date (e.g., 2023-10-01).
  - `endDate`: End date (e.g., 2023-10-31).

## Project Files

- `main.go`: Entry point of the application.
- `db/db.go`: Database connection and operations.
- `services/logger.go`: Logger service for creating and retrieving log records.
- `handlers/logger.go`: HTTP handlers for the logging API.
- `models/record.go`: Log record model.
- `schema.sql`: SQL schema for creating the log records table.

## License

This project is licensed under the MIT License.