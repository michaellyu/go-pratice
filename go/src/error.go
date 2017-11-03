package main

import (
  "fmt"
  "errors"
)

func findIndex(array []int, target int) (int, error) {
  length := len(array)
  i := 0;

  for i < length {
    if array[i] == target {
      return i, nil
    }

    i++
  }
  return -1, errors.New("Not Found")
}

func main() {
  array := []int { 1, 2, 3, 4, 5 }

  target := 3
  index, err := findIndex(array, target)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("%d index %x of %d \n", target, array, index)
  }

  target = 0
  index, err = findIndex(array, target)
  if err != nil {
    fmt.Println(err)
    fmt.Printf("%d not found in %x\n", target, array)
  } else {
    fmt.Printf("%d index %x of %d \n", target, array, index)
  }
}
