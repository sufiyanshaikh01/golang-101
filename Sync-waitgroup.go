package main

import (
	"fmt"
	"sync"
)

func Worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //signal that this goroutine is done
	fmt.Printf("Worker %d Started\n", i)

	//Some task is happening
	fmt.Printf("Worker %d End\n", i)
}

func main() {
	fmt.Println("Explore Goroutine Started")

	var wg sync.WaitGroup

	//stared in 3 worker Goroutines

	for i := 1; i <= 3; i++ {
		wg.Add(1) //Increament the WaitGroup countter
		go Worker(i, &wg)

	}

	//Wait for all Worker to finish
	wg.Wait()
	fmt.Println("worker task completed...")

}
