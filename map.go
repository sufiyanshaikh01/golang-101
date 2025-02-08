package main

import "fmt"

func main() {

	fmt.Println("We are Starting of Map in Go Lang")

	studentGrades := make(map[string]int)

	studentGrades["Sufiyan"] = 100
	studentGrades["Alin"] = 93
	studentGrades["Bob"] = 56
	studentGrades["Prince"] = 85

	fmt.Println("Marks of Bob is: ", studentGrades["Bob"])

	//change of value syntax
	studentGrades["Bob"] = 78
	fmt.Println("New Marks of Bob is: ", studentGrades["Bob"])

	//delete value syntax
	delete(studentGrades, "Bob")
	fmt.Println("...makrs is: ", studentGrades["Bob"])

	grades, exists := studentGrades["admin"]
	fmt.Println("Grades of admin: ", grades)
	fmt.Println("admin of exists: ", exists)

	Grades, Exists := studentGrades["Sufiyan"]
	fmt.Println("..Grades of Sufiyan: ", Grades)
	fmt.Println("..Sufiyan of Exists: ", Exists)

	for index, value := range studentGrades {
		fmt.Printf("key is: %s and mark is : %d\n", index, value)
	}

}
