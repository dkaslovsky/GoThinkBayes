package main

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/exercises"
)

func main() {
	fmt.Println("Cookie (manual calculation):")
	exercises.CookieManual()

	fmt.Println("Cookie:")
	exercises.Cookie()

	fmt.Println("Monty Hall:")
	exercises.MontyHall()

	fmt.Println("M&Ms:")
	exercises.MMs()

	fmt.Println("Dice:")
	exercises.Dice()
}
