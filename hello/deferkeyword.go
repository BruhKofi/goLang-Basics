package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("dealers.txt")
	// defer will run after writetofile runs or of panic runs
	d1 := dealerShip{"A1 AUTO", "MIAMI"}
	defer closeFile(file)
	writeToFile(file, d1.name)
}

type dealerShip struct {
	name string
	city string
}

func createFile(path string) *os.File {
	fmt.Println("creating file")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func writeToFile(file *os.File, dealerName string) {
	fmt.Println("writing to file")
	fmt.Fprint(file, dealerName)
}

func closeFile(file *os.File) {
	fmt.Println("closing file")
	file.Close()
}
