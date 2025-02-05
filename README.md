# Todo List Application with Go and Fiber

This is a simple Todo List CRUD API built using [Go](https://golang.org/) and the [Fiber](https://gofiber.io/) web framework. The application stores data in-memory, meaning all data will be lost once the server is stopped. This project serves as an introduction to building RESTful APIs with Go.

---

## Features

- **Create** a new todo item
- **Read** all todo items or a specific item by ID
- **Update** a todo item by ID
- **Delete** a todo item by ID

---

## Technologies Used

- **Go (Golang)** - Programming Language
- **Fiber** - Web Framework for Go

---

## Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or later)

---

## Installation


1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

3. **(Optional) Enable live reload with Air:**
   ```bash
   air
   ```
---

## Future Improvements

- Add file-based persistence using JSON.
- Integrate a database like SQLite or PostgreSQL.
- Add authentication for secure access.
- Write unit and integration tests.

## Acknowledgments

- [Go Fiber Documentation](https://docs.gofiber.io/)
- [Golang Documentation](https://golang.org/doc/)

---

Happy Coding! 🚀

