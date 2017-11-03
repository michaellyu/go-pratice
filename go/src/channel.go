package main

import "fmt"

func sendToChannel(channel chan string, value string) {
  go func() {
    channel <- value
  }()
}

func receiveFromChannel(channel chan string) string {
  return <- channel
}

func main() {
  var channel chan string = make(chan string)
  var sendValue string = "Value"
  sendToChannel(channel, sendValue)
  fmt.Printf("Send to channel: %s\n", sendValue)

  var receiveValue string = receiveFromChannel(channel)
  fmt.Printf("Receive from channel: %s\n", receiveValue)
}
