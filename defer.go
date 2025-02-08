package main

import "fmt"

func add(a, b int) int {

	return a + b
}

func main() {
	fmt.Println("We are Starting in Defer Keyword in Go language")

	fmt.Println("Starting of the  Programing")
	defer fmt.Println("Middle of the programing")
	ans := add(3, 4)
	defer fmt.Println("Ans is:", ans)
	fmt.Println("End of the Programing")
}
