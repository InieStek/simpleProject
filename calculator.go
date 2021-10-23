package main

import (
	"fmt"
	"github.com/pjaskulski/nbpapi"
	"strings"
	"time"
)

func main() {

	var b int
	fmt.Println("Co chcesz zrobić: ")
	fmt.Println("Urodziny: 1\nKalkulator: 2\nKantor: 3 ")
	fmt.Scanln(&b)
	switch b {
	case 1:
		birthday()
	case 2:
		calculator()
	case 3:
		cantor()
	}

}

func birthday() {
	var birth int
	t := time.Now()
	fmt.Println(t)
	yr := t.Year()

	fmt.Println("Kiedy się urodziłeś: ")
	var liczba int
	fmt.Scanln(&liczba)
	birth = yr - liczba
	fmt.Println(birth)

}
func calculator() {
	var wynik int
	fmt.Println("Podaj pierwsza liczbę: ")
	var first int
	fmt.Scanln(&first)

	fmt.Println("Podaj drugą liczbę: ")
	var second int
	fmt.Scanln(&second)
	fmt.Println("Jakie działanie chcesz zrobić: ")
	var cal string
	fmt.Scanln(&cal)
	if cal == "+" {
		wynik = first + second
		fmt.Println("Twój wynik to: ", wynik)
	} else if cal == "-" {
		wynik = first - second
		fmt.Println("Twój wynik to: ", wynik)

	} else if cal == "*" {
		wynik = first * second
		fmt.Println("Twój wynik to: ", wynik)

	} else if cal == "/" {
		var reszta int
		wynik = first / second
		reszta = first % second
		fmt.Println("Twój wynik to: ", wynik, "a reszta to:", reszta)

	} else {
		fmt.Println("Nie poprawne działanie")
	}
	fmt.Println("Podaj w jakich jednostkach ma być policzone działanie: ")
	var podaj string
	fmt.Scanln(&podaj)
	if podaj == "16" {
		podaj = "%x"
	} else if podaj == "8" {
		podaj = "%o"
	} else if podaj == "2" {
		podaj = "%b"
	} else if podaj == "10" {
		podaj = "%d"
	} else {
		fmt.Println("Zła dana")
	}

	h := fmt.Sprintf(podaj, wynik)

	fmt.Println(h)
}
func cantor() {
	var pln float64
	fmt.Println("Ile zapłaciłeś?")
	fmt.Scanln(&pln)
	var code string
	fmt.Println("Podaj kod waluty")
	fmt.Scanln(&code)

	value := nbpapi.NewCurrency("A")
	price, err := value.GetRateToday(strings.ToUpper(code))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(price.Mid, " "+value.Exchange.Code)
	}
	execute := price.Mid
	result := pln / execute
	fmt.Print("Zapłaciłeś w " + value.Exchange.Code + " i wyniosło to ")
	fmt.Printf("%.2f", result)
}
