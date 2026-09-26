# Todo REST API

A lightweight and fast RESTful API for managing a Todo list, built with **Go** and the **Gin Web Framework**. This project uses an in-memory data structure to store tasks, making it an excellent lightweight service or learning template for Go backend development.

## Features

* **Retrieve All Tasks:** Fetch a complete list of all current todos.
* **Retrieve a Single Task:** Fetch a specific todo using its unique ID.
* **Create Tasks:** Add new todos to the list via JSON payload.
* **Update Status:** Instantly toggle a task's completion status (from incomplete to complete, and vice versa).

## Tech Stack

* **Language:** Go (Golang)
* **Framework:** [Gin](https://gin-gonic.com/) (HTTP web framework)

## Prerequisites

* Go installed on your local machine (version 1.16+ recommended).

## Installation & Setup

1. **Clone or create the project directory:**
   ```bash
   mkdir todo-api
   cd todo-api