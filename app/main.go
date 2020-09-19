package main

import "fmt"

func main() {
	var name string = "John Doe"
	fmt.Println("This works!")
	fmt.Printf("Hello %s", name)

	amount := getAmount()
	fmt.Print(amount)
	var days int = 27
	fmt.Printf("\n%i", days)
}

func getAmount() float64 {
	return 45.555
}
