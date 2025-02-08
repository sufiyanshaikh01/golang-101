package main

import "fmt"

func SimpleFunction() {
	fmt.Println("simpole function")
}
func add(a, b int) (result int) {
	result = a + b
	return
}
func multi(a, b int) (result int) {
	result = a * b
	return
}
func main() {
	fmt.Println("Hello Go language Function Start")
	SimpleFunction()

	ans1 := add(3, 4)
	fmt.Println("addtion of two no is:", ans1)

	ans2 := multi(5, 4)
	fmt.Println("The Two no multiply is:", ans2)
}
