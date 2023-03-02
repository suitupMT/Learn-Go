package main

import "fmt"

var score = 99.5

// cannot use shorthand outside of functions
// scoreTwo := 50

func main() {
	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}

// to run you must specify  go run main.go greetings.go
//both files must be run together
