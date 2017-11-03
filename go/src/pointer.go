package main

import "fmt"

func plus(value int) {
  value++
}

func increment(value *int) {
  *value++
}

func main() {
  var value int = 0;
  fmt.Printf("value: %d\n", value)

  plus(value)
  fmt.Printf("plus(value)\n")
  fmt.Printf("value: %d\n", value)

  increment(&value)
  fmt.Printf("increment(&value)\n")
  fmt.Printf("value: %d\n", value)
}
