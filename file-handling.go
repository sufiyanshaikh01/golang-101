package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("WE are statring File handing in go lang")

	//Open file in file Handling method

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file")
		return
	}
	defer file.Close()

	//Create the buffer to read the file content
	buffer := make([]byte, 1024)

	//Read the file content into a buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file")
			return
		}

		//process the read content
		fmt.Println(string(buffer[:n]))
	}

}
