package main

import "fmt"

func sequeue(seed int) func() int {
  return func() int {
    seed += 1
    return seed
  }
}

func main() {
  var seq1 func() int = sequeue(100000)
  fmt.Println("Initial seq1(100000)")

  var seq2 func() int = sequeue(900000)
  fmt.Println("Initial seq2(900000)")

  fmt.Printf("seq1: %d\n", seq1())
  fmt.Printf("seq1: %d\n", seq1())
  fmt.Printf("seq1: %d\n", seq1())

  fmt.Printf("seq2: %d\n", seq2())
  fmt.Printf("seq2: %d\n", seq2())
  fmt.Printf("seq2: %d\n", seq2())
}
