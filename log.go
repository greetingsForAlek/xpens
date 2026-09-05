package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
		return
	}

	logExpense(name, costFloat)
}

func logExpense(name string, cost float64) {
	now := time.Now()

	year := now.Format("2006")
	month := now.Format("Jan")

	path := fmt.Sprintf("logs/%s/%s", year, month)

	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("Error creating folders:", err)
		return
	}
}