package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func logPrompt() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Item Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("Item Cost: ")
	cost, _ := reader.ReadString('\n')

	cost = strings.TrimSpace(cost)

	costFloat, err := strconv.ParseFloat(cost, 64)
	if err != nil {
		fmt.Println("Error parsing cost | Cost must be a float. (e.g. 12.99)")
	}

	logExpense(name, costFloat)
}

func logExpense(name string, cost float64) {
	fmt.Println(name, cost)
}