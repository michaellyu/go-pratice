package main

import "fmt"

func hasProperty(data map[string] string, property string) {
  value, ok := data[property]
  if (ok) {
    fmt.Printf("have a property \"%s\", and the value is \"%s\"\n", property, value)
  } else {
    fmt.Printf("have not a property \"%s\"\n", property)
  }
}

func main() {
  var news = map[string] string { "id": "1", "title": "Big News" }
  for property, value := range news {
    fmt.Printf("news[\"%s\"]: \"%s\"\n", property, value)
  }
  hasProperty(news, "title")

  delete(news, "title")
  fmt.Println("delete(news, \"title\")")
  hasProperty(news, "title")
}
