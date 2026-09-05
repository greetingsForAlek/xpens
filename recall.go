package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recallPrompt() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Year (e.g. 2026): ")
	recallYear, _ := reader.ReadString('\n')

	fmt.Println("Month (e.g. Jan, Feb, Mar, etc.): ")
	recallMonth, _ := reader.ReadString('\n')

	fmt.Println("Date (e.g. 5th = 5, 12th = 12): ")
	recallDay, _ := reader.ReadString('\n')

	recallYear = strings.TrimSpace(recallYear)
	recallMonth = strings.TrimSpace(recallMonth)
	recallDay = strings.TrimSpace(recallDay)

	readFromFile(recallYear, recallMonth, recallDay)
}

func readFromFile(year string, month string, day string) {
	path := fmt.Sprintf("logs/%s/%s/%s.txt", year, month, day)

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading selected file from path: '%s' | %s", path, err)
		return
	}

	expenses := strings.Split(string(data), "\n")
	var total float64

	fmt.Printf("\n\n\n")
	for _, expense := range expenses {
		if expense == "" {
			continue
		}
		
		parts := strings.Split(expense, ",")
		name := parts[0]
		cost := parts[1]
		fmt.Printf("Purchased %s for %s\n", name, cost)
		
		costParsed, err := strconv.ParseFloat(cost, 64)
		if err != nil {
			fmt.Println("Error parsing cost |", err)
			return
		}

		total = total + costParsed
	}
	fmt.Printf("\n\n\n")
	
	fmt.Println("Total for the day:", total)
}