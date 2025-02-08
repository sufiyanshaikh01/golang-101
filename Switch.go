package main

import "fmt"

func main() {

	fmt.Println("We are Starting of switch case for Go lang")

	fmt.Println("Enter the Two number")

	var (
		a, b int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("First number is: ", a)
	fmt.Println("second number is: ", b)

	var choice int
	fmt.Println("1. Addition")
	fmt.Println("2. subtraction")
	fmt.Println("3. multiplication")
	fmt.Println("4. division")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Addition of two number is:", a+b)
	case 2:
		fmt.Println("subtraction of two number is:", a-b)
	case 3:
		fmt.Println("multiplication of two number is:", a*b)
	case 4:
		fmt.Println(" division of two number is:", a/b)
	default:
		fmt.Println("please try again")
	}

}
