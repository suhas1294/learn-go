package main

import (
	"fmt"
	"runtime"
	"sync"
)

// variable "wait_group" from the package 'sync', wait_group is in package scope
var wait_group sync.WaitGroup

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	wait_group.Add(1) // wait for one thing, in this example its foo()
	go foo()
	bar()

	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	wait_group.Wait() // wait untill all things that are added (Add()) is done
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
	wait_group.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}

/*
Execution order in this program: (In case of a single CPU)
1. main()
2. (4) prints
3. bar()
4. (2) prints
5. foo()
6. func main() ends

Documentation:
Race condition in concurrnt programming is when, there is a shared variable and read-write get screwed up
since we are using different thread accessing that variable simultaneously
*/
