package main

import (
	"fmt"
)

func main() {
	d1 := "A1 Auto"
	d2 := "SOme other"
	d3 := "Ice Cars"

	printDealers(d1, d2, d3)
	fmt.Print("\n")
	dealers := []string{d1, d2, d2}
	printDealers(dealers...)
}

func printDealers(dealers ...string) {
	for _, dealerName := range dealers {
		fmt.Println(dealerName)
	}
}
