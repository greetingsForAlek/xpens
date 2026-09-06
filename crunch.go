package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		fmt.Println("Current Month")
	case "B":
		fmt.Println("Specific Month")
	default:
		fmt.Println("Invalid Option.")
	}
}