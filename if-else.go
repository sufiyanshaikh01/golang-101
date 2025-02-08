package main

import "fmt"

func main() {
	x := 12
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	y := 5
	if y > 5 {
		fmt.Println("y is greater than 5")
	} else if y > 10 {
		fmt.Println("y is greater than 10 but y is less than 5")
	} else {
		fmt.Println("y is less than 10")
	}
}
