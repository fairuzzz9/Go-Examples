package main

import "fmt"

func main() {
	name := "Air Asia Academy"
	//Secnario 1 -> without pointers
	fmt.Println("Outside function 1: ", name)
	printMessage(name)
	fmt.Println("Outside function 2: ", name)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Value in name: ", name)
	fmt.Println("Address of name: ", &name)
	fmt.Println("--------------------------------------------------")

	//scenario 2 -> with pointers
	fmt.Println("Outside function 1: ", name)
	printMessagePointer(&name)
	fmt.Println("Outside function 2: ", name)

}

func printMessage(message string) {
	message = "I am being spied!"
	fmt.Println("In function: ", message)
}

func printMessagePointer(message *string) {
	*message = "I am being spied!"
	fmt.Println("In function: ", *message)
}
