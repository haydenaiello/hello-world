package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println("Despacito!")

	fmt.Println(13)

	fmt.Println(13 + 1)

	fmt.Println("13+1")

	fmt.Println("Hayden Aiello")

	name := "Hayden Aiello"

	fmt.Println(name)

	var age int

	age = 13

	fmt.Println(age)

	const hairColor = "black"

	fmt.Println(hairColor)

	sayHello(name)

	despacito()
}

func sayHello(name string) {
	fmt.Print("Hello, ")
	fmt.Println(name)
}

func despacito() {
	despacito := "despacito"
	var newDespacito string

	for i, letter := range despacito {
		if i%2 != 0 {
			newDespacito = newDespacito + strings.ToUpper(string(letter))
		} else {
			newDespacito = newDespacito + string(letter)
		}
	}

	fmt.Println(newDespacito)
}
