package main

import "fmt"

// func updateName(x string) {
// 	x = "yoshi"
// }

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	//passed functions can's change the object passed into it
	name := "tifa"

	// updateName(name)
	name = updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	// pointer wrapper values
	//pointers will updates the original object without a return statement
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
