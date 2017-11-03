package main

import (
  "fmt"
  "time"
)

func runTicker(fn func (), duration time.Duration) (ticker *time.Ticker) {
  ticker = time.NewTicker(duration)
  go func () {
    for {
      _, ok := <-ticker.C
      if !ok {
        break
      }
      fn()
    }
  }()
  return ticker
}

func main() {
  i := 1
  ticker := runTicker(func () {
    fmt.Printf("Ticker running %d times\n", i)
    i++
  }, 1 * time.Second)

  time.Sleep(10 * time.Second)
  ticker.Stop()
  fmt.Println("Ticker is stoped")
}
