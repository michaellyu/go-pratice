package main

import "fmt"

func main() {
  var array [10]int
  array[0] = 1
  array[9] = 10

  fmt.Printf("array: %v\n", array)

  for i := 0; i < len(array); i++ {
    fmt.Printf("array[%d]: %d\n", i, array[i]);
  }
}
