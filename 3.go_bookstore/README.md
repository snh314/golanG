````md
# 📚 Go Bookstore Management API

A robust RESTful CRUD API built with Golang for managing a bookstore database. This project utilizes `gorilla/mux` for dynamic routing and `GORM` as the Object Relational Mapper (ORM) to seamlessly interact with a MySQL database.

## 🚀 Tech Stack

- **Language:** Go (Golang)
- **Router:** [Gorilla Mux](https://github.com/gorilla/mux)
- **ORM:** [GORM](https://gorm.io/) (v1)
- **Database:** MySQL (via XAMPP)

## 📂 Project Structure

```text
go_bookstore/
├── cmd/
│   └── main/
│       └── main.go           # Application entry point
├── pkg/
│   ├── config/               # Database connection and configuration (app.go)
│   ├── controllers/          # Business logic and request handlers (book-controller.go)
│   ├── models/               # GORM database schema and queries (book.go)
│   ├── routes/               # API endpoint definitions (bookstore-routes.go)
│   └── utils/                # Helper functions for JSON parsing (utils.go)
├── go.mod                    # Module dependencies
└── README.md
```
````

## 🛠️ Setup & Installation

1. **Prerequisites:** Ensure you have Go installed and a MySQL server running (e.g., using XAMPP).
2. **Database Setup:** Open your MySQL client (or phpMyAdmin) and create a database named `simplerest` (or match the name in your `config/app.go` file).
3. **Run the Application:**
   Navigate to the root directory and start the server:

```bash
go run cmd/main/main.go

```

_The server will start on `http://localhost:9010_`

---

## 🧪 Postman Testing Guide (CRUD Operations)

Here is how you can test all the CRUD operations using Postman. For requests that require a body, ensure you select **Body > raw > JSON** in Postman.

### 1. Create a Book (POST)

To create multiple books, send this request one by one with different data.

- **Method:** `POST`
- **URL:** `http://localhost:9010/book/`
- **JSON Body:**

```json
{
  "name": "Pather Panchali",
  "author": "Bibhutibhushan Bandyopadhyay",
  "publication": "Mitra & Ghosh"
}
```

### 2. Get Full List of Books (GET)

Retrieves all books currently stored in the database.

- **Method:** `GET`
- **URL:** `http://localhost:9010/book/`
- **JSON Body:** _None required_

### 3. Get a Single Book by ID (GET)

Retrieves the details of a specific book using its database ID.

- **Method:** `GET`
- **URL:** `http://localhost:9010/book/5` _(Replace '5' with the actual book ID)_
- **JSON Body:** _None required_

### 4. Update a Book (PUT)

Updates the information of an existing book. You must provide the ID in the URL.

- **Method:** `PUT`
- **URL:** `http://localhost:9010/book/5` _(Replace '5' with the target book ID)_
- **JSON Body:**

```json
{
  "name": "Gitanjoli (Updated Edition)",
  "author": "Rabindranath Tagore",
  "publication": "Macmillan"
}
```

### 5. Delete a Book (DELETE)

Removes a specific book from the database permanently.

- **Method:** `DELETE`
- **URL:** `http://localhost:9010/book/5` _(Replace '5' with the actual book ID)_
- **JSON Body:** _None required_

```

```
