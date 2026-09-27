# Go URL Shortener API

A simple, fast, and lightweight RESTful API for shortening URLs, built entirely with Go's standard library. It uses an in-memory data store and generates unique short links using MD5 hashing.

## Features

* **Shorten URLs:** Convert long URLs into compact, easy-to-share 8-character short links.
* **Redirection:** Automatically redirect users from a short link to the original long URL.
* **Zero Dependencies:** Built exclusively using Go's standard library (`net/http`, `crypto/md5`, etc.) without any external third-party frameworks.
* **In-Memory Storage:** Fast key-value mapping for immediate retrieval (ideal for testing and lightweight use).

## Tech Stack

* **Language:** Go (Golang)
* **Hashing Algorithm:** MD5 (First 8 characters)
* **Storage:** In-Memory Map

## Prerequisites

* Go installed on your local machine (version 1.16 or higher).

## Installation & Setup

1. **Clone or create the project directory:**
   ```bash
   mkdir url-shortener
   cd url-shortener