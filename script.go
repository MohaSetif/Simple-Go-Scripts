package main

import (
	"fmt"
	"scripts/greetings"
)

func main() {
	// count := 5
	// for i := 0; i < count; i++ {
	// 	res1 := strings.Repeat("*", i)
	// 	fmt.Println(res1)
	// }
	hi, err := greetings.Hello("Khalil")
	fmt.Println(hi)
	fmt.Print(err)
}
