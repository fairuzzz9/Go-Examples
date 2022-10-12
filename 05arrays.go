package main

import "fmt"

func main() {
	//declare array
	var scores = [4]int{88, 44, 31, 8}
	fmt.Println(scores)
	fmt.Println("Length of scores array: ", len(scores))

	//let go calculate the size of an array
	var scoresNew = [...]int{81, 71, 63, 22, 17, 26, 44, 53, 35, 88, 90, 62}
	fmt.Println(scoresNew)
	fmt.Println("Length of scoresNew array: ", len(scoresNew))

	//create a blank array
	var emptyArray = [8]int{}
	fmt.Println(emptyArray)
	fmt.Println("Length of emptyArray array: ", len(emptyArray))

	//initialize specific position with value in an array
	var specificPosition = [10]int{0: 8, 5: 22, 9: 99}
	fmt.Println(specificPosition)
	fmt.Println("Length of specificPosition array: ", len(specificPosition))

	scores[2] = 13
	fmt.Println(scores)

}
