package main

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/exercises"
)

func main() {
	fmt.Println("Cookie (by hand calculation):")
	exercises.CookieByHand()
	fmt.Println("Cookie:")
	exercises.Cookie()

	fmt.Println("Monty Hall:")
	exercises.MontyHall()

	fmt.Println("M&Ms:")
	exercises.MMs()

	fmt.Println("Dice:")
	exercises.Dice()
}
