package myPackage

import (
	"fmt"
)

func Array_test() string {

	fmt.Println("array_test function...")

	/*
			var intArray [5]int // with 'var', declaration only
			intArray := [5]int{10, 20, 30, 40, 50}
			intArray := [...]int{10, 20, 30, 40, 50, 60} // '...'
		        intArray := [5]int{1: 50, 4: 100}
	*/

	intArray := [4]*int{0: new(int), 1: new(int), 2: new(int), 3: new(int)}

	*intArray[0] = 10
	*intArray[1] = 20
	*intArray[2] = 30
	*intArray[3] = 40

	for i := 0; i < 4; i++ {
		fmt.Print("value ", *intArray[i])
	}
	fmt.Println("")

	return "returned string.."
}
