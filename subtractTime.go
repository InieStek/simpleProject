package main

import (
	"fmt"
	"time"
)

func main() {
	var years, month, day int
	currentTime := time.Now()
	fmt.Println("Podaj rok:")
	fmt.Scanln(&years)
	fmt.Println("Podaj miesiąc")
	fmt.Scanln(&month)
	fmt.Println("Podaj dzień")
	fmt.Scanln(&day)

	tmp := currentTime.AddDate(-years, -month, -day)
	fmt.Printf("Mineło: \n %d Lat\n %02d Miesięcy\n %02d Dni",
		tmp.Year(), tmp.Month(), tmp.Day())
}
