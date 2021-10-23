package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var name string
	fmt.Println("Podaj swoje imie: ")
	fmt.Scanln(&name)

	x := rand.Intn(100)
	fmt.Println(x)
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
