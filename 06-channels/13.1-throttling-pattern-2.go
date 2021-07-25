package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// defining wait group
	var wg sync.WaitGroup

	// preparing data stream
	var dataArray []int
	for i := 0; i < 100; i++ {
		dataArray = append(dataArray, i)
	}

	// defining output channel
	resultChannel := make(chan string)

	// throttling the number of goroutines -  It should be a buffered channel
	throttle := make(chan bool, 10)

	// start receiving output stream and print it
	go func() {
		for squareNumber := range resultChannel {
			fmt.Println(squareNumber)
		}
	}()

	for value := range dataArray {
		wg.Add(1)
		throttle <- true
		go returnSquare(throttle, resultChannel, &wg, value)
	}
	wg.Wait()
	close(resultChannel)
	close(throttle)
}

func returnSquare(throttle chan bool, outputChannel chan string, wg *sync.WaitGroup, input int) {
	defer wg.Done()
	outputChannel <- fmt.Sprintf("Square of %v is %v", input, input*input)
	time.Sleep(50 * time.Millisecond)
	<-throttle
}
