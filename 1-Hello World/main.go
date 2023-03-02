// entry file usually called main

package main //this is the name of the collection of code

import "fmt" //package from Go standard library

func main() { //entry point of app - if named something else will not start - just like java or C#

	fmt.Println("Hello World") //calls the package (or class but we don't call it that), and then a print line

}

//TO RUN THE PROGRAM
//Type in terminal:  go run main.go   ('go run nameOfFile.go)
