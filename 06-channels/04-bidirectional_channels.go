package main

import(
  "fmt"
)

// SEND 'TO' CHANNEL
// RECEIVE 'FROM' CHANNEL
// read from left to write for better understanding
// make(chan <- int, 2) send to channel
// make(<-chan int, 2) receive from channel

func main() {
    send_only_channel := make(chan <- int, 3) // send only channel
    fmt.Printf("typ of channel : %T", send_only_channel)

    send_only_channel <- 42
    // fmt.Println(<- send_only_channel) // Gives error since 'send_only_channel' is send only channel, cannot receive

    receive_only_channel := make(<- chan int, 3)
    // receive_only_channel <- 34 // gives error, cannot send data to a receive only channel
    //fmt.Println(<-receive_only_channel)
}