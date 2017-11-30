package main

import (
	"fmt"
	"log"

	"github.com/woocholshin/DLT/myPackage"
)

func init() {
	fmt.Println("init() here..")
}

func main() {
	fmt.Println("vim-going!!")

	/*
		myArray, err := array_test()
		if err != nil {
			log.Fatal("exit...")
		}
	*/

	myStr := myPackage.Array_test()
	fmt.Println(myStr)

	myNewStr := myPackage.Slice_test()
	fmt.Println(myNewStr)

	myPackage.Init()
	myMap := myPackage.Map_test()
	fmt.Println(myMap)

	myType, err := myPackage.Type_test()
	if err == "" {
		log.Fatal("exit...")
	}
	fmt.Println(myType, err)

	myCon := myPackage.Concurrent_test()
	fmt.Println(myCon)
}
