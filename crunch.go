package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filePath"
	"strconv"
	"strings"
	"time"
)

func crunchPrompt() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n\n\n")
	fmt.Println("What data statistics should be crunched? (Select one option)")
	fmt.Println("A - Current Month")
	fmt.Println("B - Specific Month")
	option, _ := reader.ReadString('\n')

	option = strings.TrimSpace(option)
	option = strings.ToUpper(option)

	switch (option) {
	case "A":
		currentMonth()
	case "B":
		fmt.Println("Specific Month")
	default:
		fmt.Println("Invalid Option.")
	}
}

func currentMonth() {
	now := time.Now()

	year := now.Format("2006")
	month := now.Format("Jan")

	path := fmt.Sprintf("logs/%s/%s", year, month)

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory |", err)
		return
	}

	var total float64

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(path, entry.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading", entry.Name(), ":", err)
			continue
		}

		fmt.Println("\n---", entry.Name(), "---")
		total += splitAndPrint(string(data))
	}

	fmt.Printf("\nTotal for the month: €%.2f\n", total)
}

func splitAndPrint(data string) float64 {
	lines := strings.Split(strings.TrimSpace(data), "\n")

	var monthlyTotal float64

	for _, line := range lines {
		parts := strings.Split(line, ",")

		name := parts[0]
		costString := parts[1]

		cost, err := strconv.ParseFloat(costString, 64)
		if err != nil {
			fmt.Println("Error parsing cost: ", err)
			continue
		}

		fmt.Printf("%s: €%.2f\n", name, cost)

		monthlyTotal += cost
	}

	return monthlyTotal
}