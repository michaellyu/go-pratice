package main

import "fmt"

func main() {
  var array = []int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
  for i := range array {
    fmt.Printf("array[%d]: %d\n", i, array[i])
  }

  var news = map[string] string { "id": "1", "title": "Big News" }
  for property := range news {
    fmt.Printf("news[\"%s\"]: \"%s\"\n", property, news[property])
  }

  for property, value := range news {
    fmt.Printf("news[\"%s\"]: \"%s\"\n", property, value)
  }
}
