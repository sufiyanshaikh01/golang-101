package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	ItStudent bool   `json:"its_Student"`
}

func main() {
	fmt.Println("We are Statring Json in GO language")

	person := Person{Name: "Sufiyan", Age: 19, ItStudent: true}
	fmt.Println("Person Data is: ", person)

	// convert person into Json Encoding (Marshalling)

	JsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling", err)
		return
	}
	fmt.Println("Person Data is convert to Encoding is:", string(JsonData))

	// Decoding (Unmarshalling)

	var personData Person
	err = json.Unmarshal(JsonData, &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling", err)
		return
	}
	fmt.Println("Person Data is After Deconding is", personData)

}
