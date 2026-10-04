package main

import (
	"fmt"
	"math/rand"
)

func main() {
	attempts := 0
	maxAttempts := 7

	secretNumber := rand.Intn(100) + 1

	for {
		attempts++
		fmt.Printf("Attempt %d/%d\n", attempts, maxAttempts)
		fmt.Println("================================")
		fmt.Println("       GUESSING THE NUMBER")
		fmt.Println("================================")

		var n int
		fmt.Println("Enter valid number (1-100)")
		fmt.Printf("your number: ")
		fmt.Scanln(&n)

		if n > 100 || n < 1 || n == 0 {
			fmt.Println("INVALID NUMBER")
			attempts--
		} else if n > secretNumber {
			fmt.Println("Go lower")
		} else if n < secretNumber {
			fmt.Println("Go Higher")
		} else if n == secretNumber {
			fmt.Println("Awesome! You guessed correctly.")
			fmt.Println("The secret number is:", secretNumber)
			fmt.Printf("You won in %d attempts!\n", attempts)
			break
		}

		// ম্যাক্সিমাম এটেম্পট শেষ হয়ে গেলে গেম ওভার
		if attempts == maxAttempts && n != secretNumber {
			fmt.Println("\nGame Over!")
			fmt.Printf("The number was %d.\n", secretNumber)
			break
		}
	}
}
