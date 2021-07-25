package main

import "fmt"

const (
	HTTP_NOT_FOUND = 404
	UNAVAILABLE    = "SERVICE UNAVAILABLE"
	PI_VALUE       = 3.14
	IS_SKY_BLUE    = true
)

func main() {
	fmt.Println(HTTP_NOT_FOUND)
	fmt.Println(UNAVAILABLE)
	fmt.Println(PI_VALUE)
	fmt.Println(IS_SKY_BLUE)
}
