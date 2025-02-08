package main

import "fmt"

func main() {

	fmt.Println("We are Starting Pointer in Go language")

	var num int = 12
	var ptr *int = &num

	fmt.Println("Number is: ", num)
	fmt.Println("Ptr is: ", ptr)

	//Pointer default value is nill
	var pointer *int
	if pointer == nil {
		fmt.Println("Poniter is not assigned")
	}
}
