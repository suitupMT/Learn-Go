package main //this links the files in the same package library

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello", n)
}

func showScore() {
	fmt.Println("you scored this many points: ", score)
}
