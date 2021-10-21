package main

import "fmt"

func main() {

	var tree int
	fmt.Println("Podaj ile ma mieć rzędów gwiazdek choinka")
	fmt.Scanln(&tree)

	for i := 0; i < tree; i++ {
		for j := 0; j < tree-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (2*i + 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
