package main

import "fmt"

func main() {
	welcome()
	welcomeMessage("Working with functions!!")

	// function can return values
	sum := addition(4, 4)
	fmt.Println("Addition: ", sum)
	sumsum, productproduct := additionMultiplication(4, 4)
	fmt.Println("Sumsum :", sumsum)
	fmt.Println("Productproduct :", productproduct)

	divdiv, minusminus := divisionsubstraction(26, 2)
	fmt.Println("Divdiv: ", divdiv)
	fmt.Println("Minusminus: ", minusminus)
}

func welcome() {
	fmt.Println("I am in business class function!")
}

func welcomeMessage(message string) {
	fmt.Println(message)
	fmt.Println("Length of message:", len(message))
}

func addition(n1 int, n2 int) int {
	return n1 + n2
}

func additionMultiplication(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	product := n1 * n2
	return sum, product
}

func divisionsubstraction(n1 int, n2 int) (division int, substraction int) {
	division = n1 / n2
	substraction = n1 - n2
	return

}
