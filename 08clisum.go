package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Length of Args array:", len(os.Args))
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[len(os.Args)-1])

	sum := 0
	for i := 1; i <= len(os.Args)-1; i++ {
		intNumber, _ := strconv.Atoi(os.Args[i])
		sum = intNumber + sum
	}
	fmt.Println("Sum: ", sum)
}
