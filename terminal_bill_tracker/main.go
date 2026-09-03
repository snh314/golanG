package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// helper function, for geeting input from user
func getInput(prompt string, read *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := read.ReadString('\n')

	// remove extra/unwanted space from input
	return strings.TrimSpace(input), err
}

// function for making new bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

// Show option to user
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	// print option, which is selected by user
	switch opt {
	case "a":
		fmt.Println("you choose a")
		name, _ := getInput("food name: ", reader)
		price, _ := getInput("price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("your food is ", name, "& price is:", price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println(b.name + "'s bill has been saved successfully!")
	case "t":
		fmt.Println("you choose t")
		tip, _ := getInput("tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("you give $ ", tip, "tip !! and it's added")
		promptOptions(b)
	default:
		fmt.Println("not a valid option")
		promptOptions(b)
	}
}

func main() {
	// 1. build the bill
	myBill := createBill()

	// ২. After build, show option to user again
	promptOptions(myBill)
}
