package main

import "fmt"

func main() {
	fmt.Println("We are learining array for Go language")

	var name [10]string

	name[0] = "Shaikh"
	name[1] = "Sufiyan"
	name[2] = "khan"

	fmt.Println("Name of Person is: ", name)

	//Array intialization

	var number = [8]int{1, 2, 3, 4, 5}
	fmt.Println("Numbrs is: ", number)

	//find the array length
	fmt.Println("length of array is: ", len(number))

	//find drect value of index

	fmt.Println("valu of name at 0 index is: ", name[0])
	fmt.Println("value of name at 1 index is:", name[1])
	fmt.Println("vale od name at 2 index is:", name[2])

}
