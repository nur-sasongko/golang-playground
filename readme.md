# Scripts
`go mod tidy` to install dependencies.
`go run .` to run the app.

## Basic features:
- Create a todo with form validation.
- Get all todos with pagination.
- Get todo by id
- Update a todo with form validation.
- Delete a todo.

## 🛠 Stack and Tools

- Language: **Go (Golang)**
- HTTP server: **`net/http`** (standard library)
- Routing: Basic.
- Data storage: **In-memory** first (using slices), no database yet.

---

## 🗂 Project Structure

```bash
simple-todo-api/
├── main.go
├── handler.go
└── model.go
```

- `main.go` → Start server and set routes.
- `handler.go` → Functions for each API (create, read, update, delete).
- `model.go` → Define the `Todo` struct (your data model).

---

## 🛣️ API Endpoints

| Method | URL | Description |
| --- | --- | --- |
| POST | /todos | Create a new todo |
| GET | /todos | Get all todos |
| GET | /todos/{id} | Get a todo by ID |
| PUT | /todos/{id} | Update a todo by ID |
| DELETE | /todos/{id} | Delete a todo by ID |

## 🧪 How to Test

You can use **Postman**, **cURL**, or any API client:

- `POST http://localhost:8080/todos` → with body `{ "title": "Learn Go", "done": false }`
- `GET http://localhost:8080/todos` → Get todos
- `GET http://localhost:8080/todos/1`  → Get a todo by ID
- `PUT http://localhost:8080/todos/1` → Update a todo
- `DELETE http://localhost:8080/todos/1` → Delete a tody

## 🎯 Summary

With this simple app:

- You’ll learn Go’s HTTP server, handlers, JSON encoding/decoding.
- Practice good API design (routes, methods).
- Handle simple in-memory data without DB complexity (for now).
- Add a base response to standardized response
- Implementing form validations for create and update.