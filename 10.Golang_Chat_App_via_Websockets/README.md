# Go WebSocket Chat Application

A real-time, browser-based chat application built with Go (Golang) and WebSockets. This project demonstrates how to handle concurrent client connections, broadcast messages using Go channels, and manage WebSocket upgrades using the popular `gorilla/websocket` package.

## Features

- **Real-time Messaging:** Messages are broadcast to all connected clients instantly without page reloads[cite: 20, 23].
- **Concurrency:** Utilizes Go routines and channels to safely manage multiple user connections and message routing[cite: 20, 23].
- **Template Rendering:** Uses Go's `html/template` package with `sync.Once` for efficient and thread-safe HTML rendering.
- **Customizable Port:** Server address and port can be configured via command-line flags[cite: 22].

## Tech Stack

- **Backend:** Go (Golang)
- **WebSocket Library:** [`github.com/gorilla/websocket`](https://github.com/gorilla/websocket)[cite: 20, 23]
- **Frontend:** HTML, CSS, JavaScript (jQuery)[cite: 24, 25]

## Project Structure

```text
├── main.go          # Application entry point, HTTP routing, and template handling
├── room.go          # Chat room logic, client management, and message broadcasting
├── client.go        # WebSocket connection wrapper and read/write loops
├── go.mod           # Go module definitions
├── go.sum           # Dependency checksums
└── templates/
    └── chat.html    # Frontend UI and WebSocket client logic
    └── chat.js      # Frontend UI and WebSocket client logic
```

## Prerequisites

    Go installed on your local machine (version 1.16 or higher).

## Installation & Setup

    Clone or navigate to the project directory:
        cd 10.Golang_Chat_App_via_Websockets

## Download dependencies:

    Ensure the Gorilla WebSocket package is installed:
        go mod tidy

## Run the application:

    go run .
    The server will start on port :8080 by default.

## Test the Chat:

    Open a web browser and go to http://localhost:8080.
    Open a second browser window (or incognito mode) to the same URL.
    Type a message in one window and hit "Send"—it will appear in both windows simultaneously!

## Custom Configuration

    You can change the default port by passing the -addr flag when running the server[cite: 22].
    go run . -addr=":8080"
