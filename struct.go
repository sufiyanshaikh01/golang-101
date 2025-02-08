package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println("We are Starting Structure in Go languge")

	// first method
	var person1 person
	person1.FirstName = "Sufiyan"
	person1.LastName = "Shaikh"
	person1.Age = 19
	fmt.Println("Person1 is: ", person1)

	//second method
	person2 := person{
		FirstName: "admin",
		LastName:  "root",
		Age:       20,
	}
	fmt.Println("Persron2 is: ", person2)

	//therd method help of new key
	var person3 = new(person)
	person3.FirstName = "user"
	person3.LastName = "user2"
	person3.Age = 21
	fmt.Println("Person3 is: ", person3)
}
