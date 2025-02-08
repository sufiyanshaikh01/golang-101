package main

import "fmt"

func main() {
	age := 18
	name := "SUFIYAN SHAIKH"
	height := 5.62317

	fmt.Println("Hello World")
	fmt.Println("name ", name, "age: ", age, "Height: ", height)

	fmt.Printf("name: %s\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("height: %.2f\n", height)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of height: %T\n", height)

}
