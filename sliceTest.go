package myPackage

import (
	"fmt"
)

func Slice_test() string {

	fmt.Println("=========================================")
	fmt.Println("slice_test function...")

	/*
					mySlice := make([]int, 3, 5)
					mySlice[0] = 100

				mySlice := []int{10, 20, 30, 40, 50}

			newSlice := mySlice[1:3]
			newSlice[0] = 55
			newSlice = append(newSlice, 90)
			newSlice = append(newSlice, 100)
			newSlice = append(newSlice, 200)
			newSlice = append(newSlice, 300)

		for index, value := range mySlice {
			fmt.Printf("index %d value %d\n", index, value)
		}
	*/

	mySlice := []int{10, 20, 30, 40, 50}

	newSlice := mySlice[1:3]
	newSlice[0] = 55
	newSlice = append(newSlice, 90)
	newSlice = append(newSlice, 100)
	newSlice = append(newSlice, 200)
	newSlice = append(newSlice, 300)

	fmt.Println("old slice ", mySlice, len(mySlice), cap(mySlice))
	fmt.Println("new slice ", newSlice, len(newSlice), cap(newSlice))

	return "returned from slices_test().."
}
