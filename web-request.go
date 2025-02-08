package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("We are Starting Web services in GO language")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Println("Error getting GET response:", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("Type of response is: %T\n", res)

	// Read the reponse Body

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response:", err)
		return
	}
	fmt.Println("response is:", string(data))
}
