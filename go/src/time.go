package main

import (
  "fmt"
  "time"
)

func main() {
  var now time.Time = time.Now()
  fmt.Printf("Current Time: %v\n", now)
  var format_string string = "2006-01-02 15:04:05.999"
  time_string := now.Format(format_string)
  fmt.Printf("Format Time: %s\n", time_string)
  time_val, _ := time.Parse(format_string, time_string)
  fmt.Printf("Parse Time: %v\n", time_val)
  var unix int64 = time.Now().Unix()
  fmt.Printf("Current Time: %d\n", unix)
}
