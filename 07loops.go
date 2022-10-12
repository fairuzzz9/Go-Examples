package main

import "fmt"

func main() {

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//use for loop to print the multiplication table of n
	n := 15
	for i := 1; i <= 10; i++ {
		fmt.Println(n, " x ", i, " = ", n*i)
	}

	// for loop without increment
	counter := 1
	for counter < 4 {
		fmt.Println("Counter: ", counter)
		counter++
	}

	//using range
	var scores = [4]int{88, 44, 31, 8}
	for key, value := range scores {
		fmt.Println(key, ":", value)
	}

	//use value only using range
	for _, value := range scores {
		fmt.Println(value)
	}
}
