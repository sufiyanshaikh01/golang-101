package main

import (
	"fmt"
)

func main() {

	// Basic For Loop Example
	for i := 0; i < 5; i++ {
		fmt.Println("Number is: ", i)
	}
	// break for loop example
	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 4 { 
			break
		}
	}
	//range of for loop Example

	number := []int{11, 12, 13, 14, 15}
	for index, value := range number {
		fmt.Printf("Number of index is: %d, Number of value is: %d\n", index, value)

	}

	data := "Hello world!"
	for Index, Char := range data {
		fmt.Printf("data of index is: %d, data of char is: %c\n", Index, Char)
	}
}
