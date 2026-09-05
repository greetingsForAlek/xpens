package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Too many arguments. XPens expects only 1 argument.")
	} else if len(args) < 1 {
		fmt.Println("Not enough arguments. XPens needs 1 argument.")
	}
	
	fmt.Println(args)
}