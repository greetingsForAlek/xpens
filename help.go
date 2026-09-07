package main

import "fmt"

func helpPrompt() {
	fmt.Println("Valid Commands:")
	fmt.Println("- log (Logs an expense and saves it to a file to be recalled)")
	fmt.Println("- recall (Recalls an expense from your selected day.)")
	fmt.Println("- crunch ('crunches' the data for a selected statistic.)")
}