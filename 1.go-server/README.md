# Simple Go Web Server

A basic Go web server demonstrating how to serve static files and handle GET/POST HTTP requests using the standard `net/http` package.

## Features

- Serves static HTML files (`index.html`, `form.html`) from a `static` directory.
- Handles simple `GET` requests with route validation.
- Parses and processes `POST` form data.

## Project Structure

```text
.
├── main.go
└── static/
    ├── index.html
    └── form.html
```

## Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

## How to Run

1. Clone the repository and navigate to the project directory.
2. Ensure the `static` folder exists in the same directory as `main.go` and contains the HTML files.
3. Run the server:
   ```bash
   go run main.go
   ```
4. The server will start on port `8080`.

## Endpoints

- **`GET /`** : Serves the `index.html` static website.
- **`GET /form.html`** : Serves the HTML form page.
- **`POST /form`** : Submits the form. Extracts `name` and `address` fields.
- **`GET /hello`** : Returns a simple "hello" message. (Returns 404 for non-GET requests).

## Testing API

You can easily test the endpoints using Postman or any API client:

- Send a `GET` request to `http://localhost:8080/hello`
- Send a `POST` request to `http://localhost:8080/form` with `x-www-form-urlencoded` body containing `name` and `address` keys to verify the form parsing logic without using the frontend.
