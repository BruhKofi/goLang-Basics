package main

import (
	"fmt"
)

func main() {

	carLotA := 9
	carLotB := 22

	if carLotB >= carLotA {
		fmt.Println("greater than")
	} else {
		fmt.Println("not sgreater than")
	}

	if carLotA > 15 {
		fmt.Println("sgreater than 15")
	} else if carLotA > 9 {
		fmt.Println("sgreater than 9")
	} else {
		fmt.Println("default")
	}

	switch carLotA {
	case 15:
		fmt.Println("case is 15")
	case 9, 11, 12:
		fmt.Println("is either 9, 11, 12--", carLotA)
	default:
		fmt.Println("default not needed but what the hell")
	}
}
