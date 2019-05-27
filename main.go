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

	// despacito()

	fmt.Println(makeMeme("danny devito"))
}

func sayHello(name string) {
	fmt.Print("Hello, ")
	fmt.Println(name)
}

func despacito() {
	despacito := "despacito"
	var newDespacito string

	for i, letter := range despacito {
		switch string(letter) {
		case "e":
			letter = '3'
		case "t":
			letter = '7'
		case "o":
			letter = '0'
		case "i":
			letter = '1'
		}

		if i%2 != 0 {
			newDespacito += strings.ToUpper(string(letter))
		} else {
			newDespacito += string(letter)
		}
	}

	for {
		fmt.Println(`
		░░░░░░░░░▄░░░░░░░░░░░░░░▄░░░░
		░░░░░░░░▌▒█░░░░░░░░░░░▄▀▒▌░░░
		░░░░░░░░▌▒▒█░░░░░░░░▄▀▒▒▒▐░░░
		░░░░░░░▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐░░░
		░░░░░▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐░░░
		░░░▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌░░░ 
		░░▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌░░
		░░▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐░░
		░▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌░
		░▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌░
		▀▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐░
		▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
		▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐░
		░▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌░
		░▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐░░
		░░▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌░░
		░░░░▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀░░░
		░░░░░░▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀░░░░░
		░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▀▀░░░░░░░░
		`)
		fmt.Println(newDespacito)
	}
}

// makeMeme generates a meme.
func makeMeme(s string) string {
	var newDespacito string

	for i, letter := range s {
		switch string(letter) {
		case "e":
			letter = '3'
		case "t":
			letter = '7'
		case "o":
			letter = '0'
		case "i":
			letter = '1'
		}

		if i%2 != 0 {
			newDespacito += strings.ToUpper(string(letter))
		} else {
			newDespacito += string(letter)
		}
	}

	return newDespacito
}
