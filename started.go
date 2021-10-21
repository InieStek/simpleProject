package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Scanln(&name)

	x := 36
	for i := 0; i < x; i++ {
		if (i%5 == 0) && (i%7 == 0) {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("buzz")
		}
	}
	fmt.Println("Cześć", name)

}
