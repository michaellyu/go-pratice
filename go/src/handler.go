package main

import "fmt"

type HandlerFunc func(pattern string)

func (this HandlerFunc) ServeHTTP(pattern string) {
  this(pattern)
}

type Handler interface {
  ServeHTTP(pattern string)
}

func MyHandler(pattern string) {
  fmt.Println(pattern)
}

func main() {
  var handler Handler = HandlerFunc(MyHandler)
  handler.ServeHTTP("handler")
}
