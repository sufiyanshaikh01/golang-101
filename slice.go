package main

import "fmt"

func main() {
	fmt.Println("Learn Slice for Go language ")

	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	fmt.Println("number is:", numbers)
	fmt.Printf("numbers has data type is: %T\n", numbers)
	fmt.Println("length of Numbers is:", len(numbers))

	// find the slice, length,capacity for Go language
	// length size = capacity size
	number := []int{1, 2, 3}

	fmt.Println("Slice", number)
	fmt.Println("Lengt", len(number))
	fmt.Println("Capacity", cap(number))

	// find the slice, length,capacity for Go language but use make function
	// first is length and second is capacity

	Numbers := make([]int, 3, 5)

	fmt.Println("Slice", Numbers)
	fmt.Println("Lengt", len(Numbers))
	fmt.Println("Capacity", cap(Numbers))

}
