package main

import "fmt" //if you save after importing and it isn't used it will disappear

func main() {

	//strongly typed language

	//strings - once a type is declared the variable can never be changed to another type
	var nameOne string = "text" //explicit
	var nameTwo = "Peter"       //inferred type
	var nameThree string        //only type declared

	nameThree = nameTwo
	nameFour := "Thomas"
	// go auto infers the String Type and assigns it
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//numbers-  ints and floats
	var ageOne int = 20 //explicit
	var ageTwo = 30     //inferred
	ageThree := 56      //shortened implicit

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var numOne int8 = 25 //8 bits  the bits specify the range of digits in memory, biggest is int64
	var numTwo int8 = -128
	var numThree uint8 = 254
	//look at documentation for range of different bits and numbers
	fmt.Println(numOne, numTwo, numThree)

	//floats
	var scoreOne float32 = 45.23 // the float size for decimals must be specified with float32 or float64
	var scoreTwo float64 = 8846545123134564654897894.2225
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}

//to run --  go run main.go
