# Logging Service

This is the README file for the Logging Service. This service provides endpoints to log and retrieve log messages.

## Endpoints

### 1. Create Log Entry

- **URL:** `/logs`
- **Method:** `POST`
- **Description:** Create a new log entry.
- **Request Body:**
    ```json
    {
        "level": "info",
        "message": "This is a log message",
        "timestamp": "2023-10-01T12:00:00Z"
    }
    ```
- **Response:**
    - `201 Created` on success
    - `400 Bad Request` if the request is invalid

### 2. Retrieve Log Entries

- **URL:** `/logs`
- **Method:** `GET`
- **Description:** Retrieve all log entries.
- **Response:**
    ```json
    [
        {
            "id": "1",
            "level": "info",
            "message": "This is a log message",
            "timestamp": "2023-10-01T12:00:00Z"
        },
        ...
    ]
    ```
- **Response Codes:**
    - `200 OK` on success

### 3. Retrieve Log Entry by ID

- **URL:** `/logs/{id}`
- **Method:** `GET`
- **Description:** Retrieve a log entry by its ID.
- **Response:**
    ```json
    {
        "id": "1",
        "level": "info",
        "message": "This is a log message",
        "timestamp": "2023-10-01T12:00:00Z"
    }
    ```
- **Response Codes:**
    - `200 OK` on success
    - `404 Not Found` if the log entry does not exist

### 4. Retrieve Log Entries by Time Range

- **URL:** `/logs?start={start}&end={end}`
- **Method:** `GET`
- **Description:** Retrieve log entries within a specified time range.
- **Query Parameters:**
    - `start`: Start timestamp (ISO 8601 format)
    - `end`: End timestamp (ISO 8601 format)
- **Response:**
    ```json
    [
        {
            "id": "1",
            "level": "info",
            "message": "This is a log message",
            "timestamp": "2023-10-01T12:00:00Z"
        },
        ...
    ]
    ```
- **Response Codes:**
    - `200 OK` on success
    - `400 Bad Request` if the query parameters are invalid

### 5. Delete Log Entry

- **URL:** `/logs/{id}`
- **Method:** `DELETE`
- **Description:** Delete a log entry by its ID.
- **Response Codes:**
    - `204 No Content` on success
    - `404 Not Found` if the log entry does not exist

## Running the Service

To run the logging service, use the following command:

```sh
go run main.go
```

## License

This project is licensed under the MIT License.