package main

import "fmt"

type Card struct {
	Suit   string
	Number int
}

func main() {
	message := "check me out!"
	number, boolean, charater, floater := 5, false, 'f', 3.14159
	var location *string = &message

	fmt.Println("Hello world!")
	fmt.Println(message)
	fmt.Println("Adress of the message: ", location)
	fmt.Println()
	fmt.Print("Test of printing values: ")
	fmt.Println(number, boolean, charater, floater)
	fmt.Println()

	var card Card
	card.Number = 5
	card.Suit = "Diamonds"

	fmt.Println("Your card is the ", card.Number, " of ", card.Suit)
	fmt.Println()

	var a int = 5
	var b int = 10

	fmt.Println("a= ", a, " b= ", b)
	fmt.Println("Swapping the values!")
	swap(&a, &b)

	fmt.Println("a= ", a, " b= ", b)
}

func swap(x *int, y *int) {
	var temp int = *x
	*x = *y
	*y = temp
}
