package main

import (
  "fmt"
  "time"
)

func sendToChannelLoop(channel chan int) {
  var value int = 1;
  for value <= 10 {
    fmt.Println("Sender to sleeping 1s")
    time.Sleep(1 * time.Second)
    fmt.Println("Sender is wake up")

    fmt.Println("Send value to channel: channel <-")

    channel <- value

    value++
  }
  close(channel)
}

func receiveFromChannelLoop(channel chan int) {
  for {
    fmt.Println("Receive from channel: <- channel")
    fmt.Printf("Time: %v\n", time.Now())

    val, ok := <- channel
    if ok {
      fmt.Printf("Receive value from channel: %d\n", val)
      fmt.Printf("Time: %v\n", time.Now())
    } else {
      fmt.Println("Channel is closed")
      break
    }
  }
}

func main() {
  var channel chan int = make(chan int)

  go sendToChannelLoop(channel)

  receiveFromChannelLoop(channel)
}
