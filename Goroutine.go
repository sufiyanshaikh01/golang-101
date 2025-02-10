package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Hello Sufiyan Shaikh")
}

func sayHii() {
	fmt.Println("Hii, World!!!")
	time.Sleep(2000 * time.Millisecond)

}

func main() {
	fmt.Println("We are Staring in Goroutine in Go Language")

	go sayHello()
	go sayHii()

	time.Sleep(2000 * time.Millisecond)
}
