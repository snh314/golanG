# Guess the Number (Go CLI)

A simple, interactive Command Line Interface (CLI) game written in Go. The program generates a secret random number between 1 and 100, challenging the user to guess it within exactly 7 attempts.

## Features

- **Dynamic Feedback:** Provides logical "Go Higher" or "Go lower" hints after each incorrect guess.
- **Attempt Tracking:** Enforces a strict 7-attempt limit, keeping track of the current turn and ending the game with a "Game Over" message if the limit is reached.
- **Input Validation:** Automatically detects inputs outside the 1-100 boundary (or zero/negatives). Invalid inputs trigger an "INVALID NUMBER" warning and do not consume one of the user's valuable attempts.

## Prerequisites

- Go installed on your system.

## How to Run

1. Save the code in a file named `main.go`.
2. Open your terminal or command prompt.
3. Navigate to the directory containing the file.
4. Execute the program using the following command:

   ```bash
   go run main.go
   ```

## Example Gameplay

```
Attempt 1/7
================================
       GUESSING THE NUMBER
================================
Enter valid number (1-100)
your number: 50
Go Higher
Attempt 2/7
================================
       GUESSING THE NUMBER
================================
Enter valid number (1-100)
your number: 75
Go lower
```
