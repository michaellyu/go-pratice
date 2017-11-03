package main

import "fmt"

func fibonacci(num int) int {
  if num == 1 || num == 2 {
    return 1
  }
  return fibonacci(num - 1) + fibonacci(num - 2)
}

func main() {
  n := 10
  result := fibonacci(n)
  fmt.Printf("fibonacci(%d) = %d\n", n, result)
}
