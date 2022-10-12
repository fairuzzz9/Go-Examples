package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filedata, error := ioutil.ReadFile("filesample.txt")

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	fmt.Println("File read is success!")
	fmt.Println(filedata)
	fmt.Println(string(filedata))

}
