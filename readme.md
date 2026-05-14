```markdown
# Go Web Server - No Dependencies

A production-ready REST API web server built entirely with Go's standard library, featuring Go 1.22's enhanced routing patterns. This project demonstrates how to build a complete web server with user management (CRUD operations) using only the standard library - no third-party dependencies.

## Features

- REST API for user management (Create, Read, Delete)
- In-memory data store with thread-safe access using mutexes
- Go 1.22 method-based routing with path parameters
- JSON request/response handling
- No external dependencies - only standard library
- Clean separation of concerns (models, handlers, utils)
- Proper HTTP status codes (200, 201, 204, 400, 404)
- Concurrent request handling with goroutines

## API Endpoints

| Method | Endpoint | Description | Status Codes |
|--------|----------|-------------|--------------|
| GET | /users | Retrieve all users | 200 OK |
| POST | /users | Create a new user | 201 Created, 400 Bad Request |
| GET | /users/{id} | Retrieve a specific user | 200 OK, 404 Not Found |
| DELETE | /users/{id} | Delete a user | 204 No Content, 404 Not Found |

## Project Structure

```
Go_server_no_dep/
├── main.go              # Server setup and routing
├── src/
│   ├── models.go        # Data structures (User, UserStore)
│   ├── handlers.go      # HTTP handlers (business logic)
│   └── utils.go         # JSON helper functions
```

## Prerequisites

- Go 1.22 or higher (required for method-based routing patterns)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-web-server.git
cd go-web-server
```

2. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Usage Examples

### Create a User (POST /users)

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","age":30}'
```

Response:
```json
{
  "id": "1746912345678000000",
  "name": "Alice",
  "age": 30
}
```

### Get All Users (GET /users)

```bash
curl http://localhost:8080/users
```

Response:
```json
[
  {
    "id": "1746912345678000000",
    "name": "Alice",
    "age": 30
  },
  {
    "id": "1746912345678000001",
    "name": "Bob",
    "age": 25
  }
]
```

### Get a Single User (GET /users/{id})

```bash
curl http://localhost:8080/users/1746912345678000000
```

Response:
```json
{
  "id": "1746912345678000000",
  "name": "Alice",
  "age": 30
}
```

### Delete a User (DELETE /users/{id})

```bash
curl -X DELETE http://localhost:8080/users/1746912345678000000
```

Response: 204 No Content (empty body)

### Error Handling Examples

Invalid JSON when creating a user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice"'
```

Response: 400 Bad Request - "Invalid JSON format"

User not found:
```bash
curl http://localhost:8080/users/999
```

Response: 404 Not Found - "User not found"

## Key Go Concepts Demonstrated

- Standard library `net/http` package
- Go 1.22 routing patterns with method constraints
- Path parameter extraction using `r.PathValue()`
- Struct tags for JSON marshaling/unmarshaling
- Goroutines and concurrency safety with `sync.RWMutex`
- Defer statements for resource cleanup
- Pointer receivers for methods
- Custom helper functions for JSON handling
- HTTP status codes and response writing

## Why No Dependencies?

Go 1.22 introduced powerful routing features that eliminate the need for third-party routers like `gorilla/mux` or `chi`. The standard library now supports:

- Method-based routing (`GET /users`, `POST /users`)
- Path parameters (`/users/{id}`)
- Wildcard matching

This project demonstrates that for many REST APIs, the standard library is sufficient, reducing dependencies and improving maintainability.

## License

MIT

## Author
Devraj jha
