package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `josn:"userid"`
	Id        int    `josn:"id"`
	Title     string `josn:"title"`
	Completed bool   `josn:"completed"`
}

func main() {
	fmt.Println("We are Starting CRUD in Go laguage")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error Getting is:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error is Getting Response is:", res.Status)
		return
	}

	//Method first for Json read data
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Reading is:", err)
		return
	}
	fmt.Println("Data is: ", string(data))

	//second method for Json read data
	var todo Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		fmt.Println("Error Deconding is:", err)
		return
	}
	fmt.Println("Data is:", todo)

	fmt.Println("UserId is: ", todo.UserId)
	fmt.Println("Id is: ", todo.Id)
	fmt.Println("Title is: ", todo.Title)
	fmt.Println("Completed is: ", todo.Completed)

}
