package main

import (
	"fmt"
)

func main() {
	x := 0
	//in GO we ALWAYS USE a FOR LOOP
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//for each loop
	//index is position, val is assigned range of slice nameOfSlice
	for index, val := range names {
		fmt.Printf("the value at pos %v is %v \n", index, val)
		val = "new string"
	}
	//if you get rid of the index use an _
	for _, val := range names {
		fmt.Print(val, ",")
		val = "new string" //changing the val here does nothing because val is local
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)

}
