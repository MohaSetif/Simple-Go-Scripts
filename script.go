package main

import (
	"fmt"
	"scripts/greetings"
)

func main() {
	hi, err := greetings.Hello("Khalil")
	fmt.Println(hi)
	fmt.Print(err)
}
