package main

import (
	"fmt"
	"time"
)

func name() {
	var name string
	fmt.Println("Hello, I am your Go Utility Box!")
	fmt.Println("What's your Name ? ")
	fmt.Printf("insert your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
func ageCalculation() {
	var birthYear int
	now := time.Now()
	currentYear := now.Year()
	fmt.Println("How old are you ? ")
	fmt.Printf("insert your birth year: ")
	fmt.Scanln(&birthYear)
	age := currentYear - birthYear
	fmt.Println("Your approximate age is ", age)

}
func calculator() {
	var numOne float64
	var numTwo float64
	fmt.Printf("Enter 1st numnber: ")
	fmt.Scanln(&numOne)
	fmt.Printf("Enter 2nd numnber: ")
	fmt.Scanln(&numTwo)

	fmt.Println("Choose operation: ")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	var choice int
	fmt.Printf("Insert your choice: ")
	fmt.Scanln(&choice)
	fmt.Println("Your choice: ", choice)

	if choice == 1 {
		ans := numOne + numTwo
		fmt.Println("Result is: ", ans)
	} else if choice == 2 {
		ans := numOne - numTwo
		fmt.Println("Result is: ", ans)
	} else if choice == 3 {
		ans := numOne * numTwo
		fmt.Println("Result is: ", ans)
	} else if choice == 4 {
		ans := numOne / numTwo
		fmt.Println("Result is: ", ans)
	} else {
		fmt.Println("Insert correct number")
	}

}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func FahrenheitTocelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func tempratureConverter() {
	var userChoice int
	fmt.Println("Temprature Converter")
	fmt.Println("1. Celsius → Fahrenheit")
	fmt.Println("2. Fahrenheit → Celsius")
	fmt.Printf("Select Option: ")
	fmt.Scan(&userChoice)

	if userChoice == 1 {
		var c float64
		fmt.Printf("Enter Celsius: ")
		fmt.Scan(&c)

		ans := celsiusToFahrenheit(c)
		fmt.Println("Result in Fahrenheit:", ans)

	} else if userChoice == 2 {
		var f float64
		fmt.Printf("Enter Fahrenheit: ")
		fmt.Scan(&f)

		ans := FahrenheitTocelsius(f)
		fmt.Println("Result in Celsius:", ans)

	} else {
		fmt.Println("Enter valid number")
	}
}
func total() {
	var value int
	fmt.Println("GO UTILITY BOX")
	fmt.Println("1. Name")
	fmt.Println("2. Age Calculator")
	fmt.Println("3. Calculator")
	fmt.Println("4. Temprature Converter")
	fmt.Printf("Enter your value: ")
	fmt.Scanln(&value)
	if value == 1 {
		name()
	} else if value == 2 {
		ageCalculation()
	} else if value == 3 {
		calculator()
	} else if value == 4 {
		tempratureConverter()
	} else {
		fmt.Println("Enter valid number")
	}

}

func main() {
	total()

}
