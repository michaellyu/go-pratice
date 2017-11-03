package main

import "fmt"

type News struct {
  id int
  title string
}

func main() {
  var news News
  news.id = 1
  news.title = "Big News"
  fmt.Printf("News: { id: \"%d\", title: \"%s\" }\n", news.id, news.title)
}
