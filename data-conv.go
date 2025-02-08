package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("We are Starting data converting in go language")

	var num int = 144
	fmt.Println("Number is: ", num)
	fmt.Printf("Type of Number is: %T\n", num)

	var data float64 = float64(num)
	fmt.Println("Data is: ", data)
	fmt.Printf("Type of data is: %T\n", data)

	num = 144
	Str := strconv.Itoa(num)
	fmt.Println("String is: ", Str)
	fmt.Printf("Type of String is: %T\n", Str)

	number_Str := "12345"
	number_int, _ := strconv.Atoi(number_Str)
	fmt.Println("Number of integer is: ", number_int)
	fmt.Printf("Type of number_int is: %T\n", number_int)

	num_str := "11.43"
	num_float, _ := strconv.ParseFloat(num_str, 64)
	fmt.Println("Number flaot is: ", num_float)
	fmt.Printf("Type of num_flaot is: %T\n", num_float)

}
