package main

import "fmt"

func main() {

	counter := 20
	//if condition
	if counter > 20 {
		fmt.Println("Counter greater than 20")
	}

	//if else
	if counter > 20 {
		fmt.Println("Counter greater than 20")
	} else {
		fmt.Println("Counter not greater than 20")
	}

	//if else if
	if counter > 20 && counter < 100 {
		fmt.Println("Counter greater than 20")
	} else if counter < 0 {
		fmt.Println("Negative values are not allowed!")
	} else if counter == 20 {
		fmt.Println("Yay... you win the lottery")
	} else {
		fmt.Println("You are at 100 or more....")
	}
}
