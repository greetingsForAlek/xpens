package main

import (
	"fmt"
	"os"
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

	switch command {
	case "log":
		logPrompt()
	case "recall":
		recallPrompt()
	case "crunch":
		fmt.Println("Crunch expenses statistics")
	case "help":
		fmt.Println("Help command")
	default:
		fmt.Println("Unknown command: ", command, "| To see all commands, run 'xpens help'")
	}
}