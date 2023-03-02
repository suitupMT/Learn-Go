package main

import (
	"fmt"
	"strings"
)

//this is new

// in the signature value types are specified like a second set of parameters to indicate multiple returns
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ") //returns a slice with multiple elements

	var initials []string

	for _, v := range names { //names is an array with 2 seperate names, 1 in each element
		initials = append(initials, v[:1]) //this just grabs the range of a string 1st element
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	//if this has failed or there is no second name the it returns the first initial and then blank
	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("barret")
	fmt.Println(fn3, sn3)
}
