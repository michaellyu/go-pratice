package main

import (
  "fmt"
  "strconv"
)

func handleError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func main() {
  i, err := strconv.Atoi("10")
  handleError(err)
  fmt.Printf("i: %d\n", i)

  a := strconv.Itoa(i * i)
  fmt.Printf("a: \"%s\"\n", a)

  b, err := strconv.ParseBool("true")
  handleError(err)
  fmt.Printf("b: %v\n", b)

  f, err := strconv.ParseFloat("1.1", 0)
  handleError(err)
  fmt.Printf("f: %f\n", f)

  i64, err := strconv.ParseInt("-1", 10, 0)
  handleError(err)
  fmt.Printf("i64: %d\n", i64)

  u64, err := strconv.ParseUint("1", 10, 0)
  handleError(err)
  fmt.Printf("u64: %d\n", u64)

  s := strconv.Quote("Hello, World!")
  fmt.Printf("s: \"%s\"\n", s)

  bytes := []byte("bool:")
  bytes = strconv.AppendBool(bytes, true)
  fmt.Println(string(bytes))
}
