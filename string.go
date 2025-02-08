package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("We are starting String package in Go language")

	data := "orange,apple,mango,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three two four two five six two"
	count := strings.Count(str, "two")
	fmt.Println("Count is: ", count)

	str = "     Hello GO!       "
	trim := strings.TrimSpace(str)
	fmt.Println("Trimed is: ", trim)

	str1 := "Sufiyan"
	str2 := "Shaikh"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Result is: ", result)

}
