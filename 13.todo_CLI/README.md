# Go CLI Todo List

A lightweight, interactive Command Line Interface (CLI) Todo List application built with Go (Golang). This program allows users to seamlessly manage their daily tasks, track completion status, and organize their workflow directly from the terminal.

## Features

- **Add Tasks:** Create new tasks to keep track of your goals.
- **View Tasks:** Display all current tasks with visual status indicators (`[ ]` for pending, `[✓]` for completed).
- **Complete Tasks:** Mark specific tasks as finished using their unique ID.
- **Delete Tasks:** Permanently remove tasks from your list to keep it clean.
- **Continuous Loop:** The application runs in an interactive menu loop until you explicitly choose to exit.

## Prerequisites

- Go installed on your local machine (version 1.16 or higher is recommended).

## How to Run

1. Save the provided code in a file named `main.go`.
2. Open your terminal or command prompt.
3. Navigate to the folder containing the `main.go` file.
4. Run the application using the following command:

   ```bash
   go run main.go
   ```

## Usage Example

```
=== My Todo List ===
1. Add - Add new Todo in the list
2. View - see full Todo list
3. Complete - complete Todo list
4. Delete - Erase Todo
5. Exit - Close program

 Select an Option (1-5): 1
Add Todo Title: Coding
✓ 'Coding' is added!

 Select an Option (1-5): 2

=== Your Todos ===
1. [ ] Coding
```
