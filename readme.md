# Go Web Server (Standard Library Only)

A lightweight REST API server built entirely with Go’s standard library using Go 1.22 routing features.

This project demonstrates how to build a clean and production-ready web server without third-party dependencies while supporting user CRUD operations, JSON handling, and concurrent request safety.

````markdown
---

## Features

- REST API for user management
- Create, read, and delete users
- Thread-safe in-memory datastore using `sync.RWMutex`
- Go 1.22 method-based routing
- Path parameters with `r.PathValue()`
- JSON request and response handling
- Proper HTTP status codes
- Clean project structure
- No external dependencies

---

## API Endpoints

| Method | Endpoint      | Description         |
| ------ | ------------- | ------------------- |
| GET    | `/users`      | Get all users       |
| POST   | `/users`      | Create a user       |
| GET    | `/users/{id}` | Get user by ID      |
| DELETE | `/users/{id}` | Delete user by ID   |

### Status Codes

- `200 OK`
- `201 Created`
- `204 No Content`
- `400 Bad Request`
- `404 Not Found`

---

## Project Structure

```text
Go_server_no_dep/
├── main.go
└── src/
    ├── models.go
    ├── handlers.go
    └── utils.go
````

| File          | Purpose                  |
| ------------- | ------------------------ |
| `main.go`     | Server setup and routing |
| `models.go`   | Data models and storage  |
| `handlers.go` | HTTP handlers            |
| `utils.go`    | JSON helper functions    |

---

## Requirements

* Go 1.22 or later

---

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-web-server.git
cd go-web-server
```

### Run the Server

```bash
go run main.go
```

Server runs at:

```text
http://localhost:8080
```

---

## Usage

### Create User

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

---

### Get All Users

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
  }
]
```

---

### Get User By ID

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

---

### Delete User

```bash
curl -X DELETE http://localhost:8080/users/1746912345678000000
```

Response:

```text
204 No Content
```

---

## Error Examples

### Invalid JSON

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice"'
```

Response:

```text
400 Bad Request
```

---

### User Not Found

```bash
curl http://localhost:8080/users/999
```

Response:

```text
404 Not Found
```

---

## Go Concepts Used

* `net/http`
* Go 1.22 routing patterns
* `r.PathValue()`
* JSON encoding and decoding
* Struct tags
* `sync.RWMutex`
* Goroutines
* Pointer receivers
* Helper functions
* HTTP status handling

---

## Why No Dependencies?

Go 1.22 introduced built-in routing improvements that make external routers unnecessary for many APIs.

Supported features include:

* Method-based routing
* Path parameters
* Wildcard matching

Using only the standard library keeps the project lightweight, maintainable, and dependency-free.

---

## License

MIT

## Author


```
Devraj Jha
```
