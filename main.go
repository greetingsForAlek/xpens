package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Too many arguments. XPens expects only 1 argument.")
		return
	} else if len(args) < 1 {
		fmt.Println("Not enough arguments. XPens needs 1 argument.")
		return
	}

	fmt.Println(args)
	command := args[0]
	command = strings.ToLower(command)

	switch command {
	case "log":
		logPrompt()
	case "recall":
		recallPrompt()
	case "crunch":
		crunchPrompt()
	case "help":
		helpPrompt()
	default:
		fmt.Println("Unknown command: ", command, "| To see all commands, run 'xpens help'")
	}
}