package main

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/exercises"
)

func main() {
	fmt.Println("manual calculation:")
	exercises.CookieManual()

	fmt.Println("encapsulated calculation:")
	exercises.CookieEncapsulated()

	fmt.Println("suite calculation:")
	exercises.CookieSuite()
}
