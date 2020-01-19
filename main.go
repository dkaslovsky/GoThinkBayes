package main

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/exercises"
)

func main() {
	fmt.Println("Cookie: manual calculation:")
	exercises.CookieManual()
	fmt.Println("Cookie: suite calculation:")
	exercises.CookieSuite()

	fmt.Println("Monty Hall:")
	exercises.MontyHall()
}
