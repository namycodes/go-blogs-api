# 📝 Blogs API - Clean Architecture Backend

A simple and extensible RESTful API for managing blog posts, built in **Go** using the **Gin web framework**, **GORM** for ORM, and organized using the principles of **Clean Architecture**.

---

## 📆 Features

* ✅ Create and retrieve blog posts
* 🧱 Clean Architecture (separation of concerns)
* 🖠 GORM + PostgreSQL for persistence
* ⚡ Fast and lightweight using Gin
* 🧪 Easily testable and scalable

---

## 📂 Project Structure

```
blogs-api/
│
├── config/           # Database connection config
├── migrations/       # Auto-migrations for models
├── models/           # Database models (e.g. Blog)
├── repository/       # Data access logic (interface + implementation)
├── service/          # Business logic layer
├── handler/          # HTTP handlers (Gin controllers)
├── main.go           # App entry point
|__ internal/         # internal Workings of the app
|__ pkg/              # Database configs
└── go.mod            # Go module definition
```

---

## 🧱 Clean Architecture Layers

| Layer          | Responsibility                          |
| -------------- | --------------------------------------- |
| **Models**     | Core entities (e.g., `Blog`)            |
| **Repository** | Data access abstraction (DB read/write) |
| **Service**    | Business rules / logic                  |
| **Handler**    | HTTP request/response & validation      |

---

## 🔧 Tech Stack

* **Go (Golang)**
* **Gin** - Web framework
* **GORM** - ORM for PostgreSQL
* **PostgreSQL** - Database

---

## ✨ Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/namycodes/blogs-api.git
cd blogs-api
```

### 2. Setup your PostgreSQL database

```sql
CREATE DATABASE blogs;
```

> Make sure the connection details match your setup in `config/database.go`.

### 3. Run the app

```bash
go run main.go
```

The server will start on:
📍 `http://localhost:8080`

---

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/v1/blogs`

| Method | Endpoint | Description       |
| ------ | -------- | ----------------- |
| `GET`  | `/`      | Get all blogs     |
| `POST` | `/`      | Create a new blog |
| `GET`  | `/:id`   | Get blog by ID    |
| `DELETE`  | `/:id`   | Delete blog by ID    |
| `PATCH`  | `/:id`   | Edit blog by ID    |

Base URL: `http://localhost:8080/api/v1/auth`

| Method | Endpoint | Description       |
| ------ | -------- | ----------------- |
| `POST`  | `/login`      | User Login     |
| `POST` | `/register`      | Register New User |

---

## 🧪 Example Blog Payload

```json
{
  "title": "Intro to Clean Architecture",
  "description": "Any Description here"
}
```

---

## 🧠 Dependencies

Make sure you have the following installed:

* Go 1.18+
* PostgreSQL
* Git

---

## 🧼 Optional: .env Setup

You can manage credentials via `.env` and load them using libraries like `godotenv`.

```env
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=blogs
DB_PORT=5432
```

---

## 🧑‍💻 Author
@namycodes
[LinkedIn](https://www.linkedin.com/in/namycodes/) • [GitHub](https://github.com/namycodes)
