# Terminal Bill Tracker

A lightweight, interactive Command Line Interface (CLI) application built with Go for managing restaurant bills. This tool allows users to instantly create bills, add food items, apply tips, and generate formatted text receipts.

## Features

- **Interactive Prompts:** Step-by-step CLI interface using Go's `bufio` and standard input.
- **Dynamic Item Entry:** Add multiple food items and their respective prices to a single bill.
- **Custom Tips:** Apply tip amounts that automatically update the final total.
- **Smart File Saving:** Generates formatted receipts in a dedicated `bills/` directory. Automatically handles duplicate names (e.g., saves as `bob_2.txt` if `bob.txt` exists) to prevent overwriting.

## Project Structure

- `main.go` - Handles the application's core logic, terminal execution flow, user input, and the interactive `switch` statement menu.
- `bill.go` - Contains the `bill` struct, receiver functions (methods) for updating items and tips, formatting the receipt, and the file-saving logic.
- `bills/` - The destination folder where all generated text receipts are stored.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.15 or higher) installed on your machine.

## How to Run

1. Navigate to the project folder in your terminal.
2. Run both Go files together:
   ```bash
   go run main.go bill.go
   ```

---

## Usage Example

## Usage Example

```
Create a new bill name: bob
Created the bill - bob
Choose option (a - add item, s - save bill, t - add tip): a
food name: Burger
price: 5.99
your food is Burger & price is: 5.99
Choose option (a - add item, s - save bill, t - add tip): t
tip amount: 2.00
you give $ 2.00 tip !! and it's added
Choose option (a - add item, s - save bill, t - add tip): s
Bill is saved
bob's bill has been saved successfully!

```
