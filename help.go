package main

import (
	"fmt"
	"strings"
)

func helpPrompt() {
	fmt.Println("Valid Commands:")
	fmt.Println("- log (Logs an expense and saves it to a file to be recalled)")
	fmt.Println("- recall (Recalls an expense from your selected day.)")
	fmt.Println("- crunch ('crunches' the data for a selected statistic.)")
}

func checkCommand(command string) {
	letters := strings.Split(command, "")
	fmt.Println(letters)

	if len(letters) == 3 {
		fmt.Println("3 letters")
	} else if len(letters) == 4 {
		fmt.Println("4 letters")
	} else if len(letters) == 6 {
		fmt.Println("6 letters")
	} else {
		fmt.Println("Amount of letters does not match any command")
	}
}