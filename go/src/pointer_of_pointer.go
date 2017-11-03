package main

import "fmt"

func main() {
  var val int = 1
  var ptr *int
  var ptr_of_ptr **int
  fmt.Printf("val: %d\n", val)
  ptr = &val
  fmt.Println("ptr = &val")
  fmt.Printf("*ptr: %d\n", *ptr)
  ptr_of_ptr = &ptr
  fmt.Println("ptr_of_ptr = &ptr")
  fmt.Printf("**ptr_of_ptr: %d\n", **ptr_of_ptr)
}
