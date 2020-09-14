package main

import "fmt"

func main() {
	var name string = "John Doe"
	fmt.Println("This works!")
	fmt.Printf("Hello %s", name)

	amount := getAmount()
	fmt.Print(amount)
}

func getAmount() float64 {
	return 45.555
}
