package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("NUMBER_OF_PROCESSORS"))
	fmt.Println(os.Getenv("OS"))
	fmt.Println(os.Getenv("OBB"))

	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumCPU())

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.Day()) //gives date
	fmt.Println(currentTime.Weekday())
	fmt.Println(currentTime.Month())

}
