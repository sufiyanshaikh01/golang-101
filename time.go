package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("We are Strarting in time package in Go languag")

	currentTime := time.Now()
	fmt.Println("Curent Time is: ", currentTime)
	fmt.Printf("Type of current time is: %T\n", currentTime)

	formatted := currentTime.Format("02/01/2006, Monday, 03:04:05 PM")
	println("Time is ", formatted)

	layout := "02-01-2006"
	data_str := "29-12-2024"
	formatted_time, _ := time.Parse(layout, data_str)
	fmt.Println("Formatted_time is:", formatted_time)

	// add one more day in current Time
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new date is:", new_date)
	formatted_new_date := new_date.Format("02/01/2006, Monday")
	fmt.Println("Formatted new date is:", formatted_new_date)

}
