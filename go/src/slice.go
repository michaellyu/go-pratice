package main

import "fmt"

func printSlice(slice []int) {
  fmt.Printf("len(slice): %d, cap(slice): %d, slice: %v\n", len(slice), cap(slice), slice)
}

func main() {
  var slice []int
  slice = make([]int, 3, 5)
  printSlice(slice)

  slice[0] = 1
  slice[1] = 2
  slice[2] = 3
  slice = append(slice, 4)
  fmt.Println("slice = append(slice, 4)")
  printSlice(slice)

  slice = append(slice, 5)
  fmt.Println("slice = append(slice, 5)")
  printSlice(slice)

  slice = append(slice, 6)
  fmt.Println("slice = append(slice, 6)")
  printSlice(slice)

  var array = []int { 9, 8, 7, 6, 5, 4, 3, 2, 1, 0 }
  fmt.Printf("array: %v\n", array)
  fmt.Printf("copy(slice, array)\n")
  copy(slice, array)
  printSlice(slice)

  fmt.Printf("slice[1:2]: %v\n", slice[1:2])
  fmt.Printf("slice[1:4]: %v\n", slice[1:4])
}
