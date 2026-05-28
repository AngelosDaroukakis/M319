package main

import "fmt"

func changeNumber(number *int) {
	*number = 100
}

func main() {
	x := 42

	fmt.Println("Wert von x:", x)
	fmt.Println("Adresse von x:", &x)

	pointer := &x

	fmt.Println("Adresse im Pointer:", pointer)
	fmt.Println("Wert über Pointer:", *pointer)

	*pointer = 50
	fmt.Println("Neuer Wert von x:", x)

	changeNumber(&x)
	fmt.Println("Wert von x nach Funktion:", x)
}