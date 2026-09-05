package main

import (
	"bufio"
	"fmt"
	"os"
)

func recallPrompt() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Year (e.g. 2026): ")
	recallYear, _ := reader.ReadString('\n')

	fmt.Println("Month (e.g. Jan, Feb, Mar, etc.): ")
	recallMonth, _ := reader.ReadString('\n')

	fmt.Println("Date (e.g. 5th = 5, 12th = 12): ")
	recallDay, _ := reader.ReadString('\n')

	fmt.Println(recallYear, recallMonth, recallDay)
}