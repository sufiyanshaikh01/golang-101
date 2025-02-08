package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("We are Starting in file handilg in Go language")
	
   // Creating file in file handing metod
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()

	content := "Hello world by Sufiyan"
	byte, error := io.WriteString(file, content)
	fmt.Println("byte written is: ", byte)
	if error != nil {
		fmt.Println("Error while writing file: ", error)
		return
	}

	fmt.Println("Successfully create file")

}
