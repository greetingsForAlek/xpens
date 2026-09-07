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
		if letters[0] == "l" || letters[1] == "o" || letters[2] == "g" {
			fmt.Printf("Unrecognised command '%s'. Similar command: log\n", command)
			fmt.Println("For a list of all commands, run 'xpens help'.")
		} else {
			fmt.Printf("Unrecognised command '%s'.\n", command)
			fmt.Println("For a list of all commands, run 'xpens help'.")
		}
	} else if len(letters) == 4 {
		if letters[0] == "h" || letters[1] == "e" || letters[2] == "l" || letters[3] == "p" {
			fmt.Printf("Unrecognised command '%s'. Similar command: help. Run help for a list of commands.\n", command)
		} else {
			fmt.Printf("Unrecognised command '%s'.\n", command)
			fmt.Println("For a list of all commands, run 'xpens help'.")
		}
	} else if len(letters) == 6 {
		if letters[0] == "r" || letters[1] == "e" || letters[2] == "c" || letters[3] == "a" || letters[4] == "l" || letters[5] == "l" {
			fmt.Printf("Unrecognised command '%s'. Similar command: recall\n", command)
			fmt.Println("For a list of all commands, run 'xpens help'.")
		} else if letters[0] == "c" || letters[1] == "r" || letters[2] == "u" || letters[3] == "n" || letters[4] == "c" || letters[5] == "h" {
			fmt.Printf("Unrecognised command '%s'. Similar command: crunch\n", command)
			fmt.Println("For a list of all commands run 'xpens help'.")
		} else {
			fmt.Printf("Unrecognised command '%s'.\n", command)
			fmt.Println("For a list of all commands, run 'xpens help'.")
		}
	} else {
		fmt.Printf("Unrecognised command '%s'.\n", command)
		fmt.Println("For a list of all commands, run 'xpens help'.")
	}
}