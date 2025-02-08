package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	Email string
	Phone int
}
type Address struct {
	HouseNo int
	Area    string
	State   string
}
type Employee struct {
	Person_Detail  person
	Person_Contact Contact
	Person_Addres  Address
}

func main() {
	fmt.Println("We  are Starting second tyoe of Structure")

	var employee Employee
	employee.Person_Detail = person{
		FirstName: "Sufiyan ",
		LastName:  "Shaikh",
		Age:       19,
	}

	employee.Person_Contact.Email = "abcd@gmail.com"
	employee.Person_Contact.Phone = 7977325018

	employee.Person_Addres = Address{
		HouseNo: 14,
		Area:    "Kalyan",
		State:   "Maharastra",
	}

	fmt.Println("Employee is: ", employee)
}
