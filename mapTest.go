package myPackage

import (
	"fmt"
)

func Init() {
	fmt.Println("-----------------------------------------")
	fmt.Println("mapTest init() called..")
}

func Map_test() string {

	fmt.Println("map_test function...")

	/*
			myMap := make(map[string]int)  // string key, int value example

			//////////////////////////////////////
		myColors := map[string]string{
			"AliceBlue": "#f0f8ff",
			"Coral":     "#ff7f50",     // comma required right after the final map entity
		}

		for key, value := range myColors {
			fmt.Printf("Key %s Value %s", key, value)
		}

	*/

	colors := map[string]string{}
	colors["Red"] = "#da1337"

	value, exists := colors["Red"]

	if exists {
		fmt.Println(value)
	}

	fmt.Println("map", colors)

	myColors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral":     "#ff7f50",
	}

	for key, value := range myColors {
		fmt.Printf("[Key %s Value %s]", key, value)
	}

	delete(myColors, "Coral")

	fmt.Printf("\n-----------------------------------------\n")

	for key, value := range myColors {
		fmt.Printf("[Key %s Value %s]", key, value)
	}

	return "returned from map_test().."
}
