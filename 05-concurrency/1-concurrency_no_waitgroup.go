/*
	concurrency vs parallelism:
	Concurrency is a design pattern -  they way we design/write the code - It does not gurantee parallelism.
	When there is a Single CPU (single core), Go wont run code in parallel.
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	go foo()
	bar()

	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}

/*
Here, in this program, execution comes to `go foo()`,
it launches a new go routine and continues execution without waiting
before foo() is run, bar() and other two print statements executed in micro seconds and hence
execution ends even before foo() runs. Hence we use something called as wait groups which does the
job of synchronisation
*/
