package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println(expenses)
}