package main

import "fmt"

func main() {

	// single line commen
	/*
		multi line comment
		multi line comment
	*/

	// declare a variable
	// var variableName datatype
	var name string
	name = "Air Asia Academy"
	fmt.Println(name)
	fmt.Println(name)

	// use format specifier to check the datatype
	// wiht format specifier, need to use Printf
	fmt.Printf("I learn at %v in KL or online. \n", name)
	fmt.Printf("Data type of name is %T\n", name)
	// if you use Printf, add the new line explicitly
	// combine two format specifiers together
	fmt.Printf("Datatype of %v (name variable) is %T\n", name, name)

	// use auto type inference to declare the variable
	location := "KL Sentral"
	fmt.Println(location)
	fmt.Printf("Variable (location) value: %v Datatype: %T", location, location)

	score := 22
	points := 22.22
	check := true

	fmt.Printf("\nVariable (score) value: %v Datatype: %T\n", score, score)
	fmt.Printf("Variable (points) value: %v Datatype: %T\n", points, points)
	fmt.Printf("Variable (check) value: %v Datatype: %T\n", check, check)

}
