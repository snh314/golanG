package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// constructor function - for making new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Pointer receiver method for tip update
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add new food item in Pointer receiver method
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Method for create bill format
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// loop over items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// calculate total bil [total bill = bill + tip]
	total += b.tip
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs

}

// save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Bill is saved")

}
