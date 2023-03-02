package main

//this is how you import multiple "classes"
//notice no commas or semi-colons
import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	//Strings class
	//Contains is searching a string, for a value and returns a boolean
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.Contains(greeting, "bye"))
	//selects the string and returns a new value but does not change the variable
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	//places each word in an array between spaces
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	//create the ages slice
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names) // will actually alter the orginal slice
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser")) //returns the position in the slice

}
