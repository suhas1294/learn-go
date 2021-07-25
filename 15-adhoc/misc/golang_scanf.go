package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 Three ways of taking input
   1. fmt.Scanln(&input)
   2. reader.ReadString()
   3. scanner.Scan()

   Here we recommend using bufio.NewScanner
*/

func main() {
	// To create dynamic array
	// arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your email address: ")
	scanner.Scan()
	text := scanner.Text()

	fmt.Println(text)
	/*
		for {
			fmt.Print("Enter Text: ")
			// Scans a line from Stdin(Console)
			scanner.Scan()
			// Holds the string that scanned
			text := scanner.Text()
			if len(text) != 0 {
				fmt.Println(text)
				arr = append(arr, text)
			} else {
				break
			}

		}
		// Use collected inputs
		fmt.Println(arr)
	*/
}
