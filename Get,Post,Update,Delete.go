package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `josn:"userid"`
	Id        int    `josn:"id"`
	Title     string `josn:"title"`
	Completed bool   `josn:"completed"`
}

func performGetRequest() {
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
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error Reading is:", err)
	// 	return
	// }
	// fmt.Println("Data is: ", string(data))

	//second method for Json read data
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
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

func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "SUFIYAN SHAIKH",
		Completed: true,
	}

	//Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshalling ", err)
		return
	}

	//Convert Json data to String
	jsonString := string(jsonData)

	//convert String data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	//send POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Response: ", err)
		return
	}
	fmt.Println("Response : ", string(data))

	fmt.Println("Response status is: ", res.Status)
}

func performUpdateRequest() {
	todo := Todo{
		UserId:    291205,
		Title:     "SUFIYAN SHAIKH GOlang Hello World!",
		Completed: false,
	}

	//Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshalling ", err)
		return
	}

	//Convert Json data to String
	jsonString := string(jsonData)

	//convert String data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT Request

	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error Creating PUT Requset: ", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending Request: ", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Resopnse: ", err)
		return
	}
	fmt.Println("Response: ", string(data))

	fmt.Println("Response status is: ", res.Status)

}

func performDeleteRequest() {
	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	//create Delete Request

	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error Creating Delete Requset: ", err)
		return
	}

	//Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending Request: ", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Resopnse: ", err)
		return
	}
	fmt.Println("Response: ", string(data))

	fmt.Println("Response status is: ", res.Status)

}

func main() {
	fmt.Println("We are Starting POST Mthod in Go laguage")
	performGetRequest()
	performPostRequest()
	performUpdateRequest()
	performDeleteRequest()
}
