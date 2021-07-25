package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// step-1 preparing data
	var dataArray []int
	for i := 1; i <= 100; i++ {
		dataArray = append(dataArray, i)
	}

	// throttling
	numOfGoRoutines := 10
	wg.Add(numOfGoRoutines)

	// declaring channels
	inputChannel := make(chan int)
	resultChannel := make(chan int)

	// streaming data to input channel from data array
	go func() {
		for data := range dataArray {
			fmt.Printf("sending %v for input channel from data stream\n", data)
			inputChannel <- data
		}
		close(inputChannel)
	}()

	for j := 0; j < numOfGoRoutines; j++ {
		// 10 goroutines are launched here.
		go func() {
			// what each (out of 10) goroutine will do :
			for v := range inputChannel {
				resultChannel <- someTimeConsumingTask(v)
			}
			wg.Done()
		}()
	}
	go func() {
		for output := range resultChannel {
			fmt.Printf("received processed output %v\n", output)
		}
		close(resultChannel)
	}()
	wg.Wait()

	fmt.Println("\n\n", dataArray)
}

func someTimeConsumingTask(input int) int {
	fmt.Printf("processing %v \n", input)
	time.Sleep(100 * time.Millisecond)
	return input + rand.Intn(1000)
}
