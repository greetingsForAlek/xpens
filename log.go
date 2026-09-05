package main

import (
	"bufio"
	"fmt"
	"os"
)

func logPrompt() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Item Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("Item Cost: ")
	cost, _ := reader.ReadString('\n')

	fmt.Println(name, cost)
}