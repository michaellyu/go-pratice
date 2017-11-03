package main

import (
  "io"
  "net/http"
  "html/template"
  "log"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
  io.WriteString(response, "Hello, World\n")
}

func templateHandler(response http.ResponseWriter, request *http.Request) {
  t := template.New("template")
  t, _ = t.Parse(`{{.}} template`)
  t.Execute(response, "string")
}

func webHandler(response http.ResponseWriter, request *http.Request) {
  t, _ := template.ParseFiles("web.html")
  t.Execute(response, "file")
}

func main() {
  http.HandleFunc("/", helloHandler)
  http.HandleFunc("/template", templateHandler)
  http.HandleFunc("/web", webHandler)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
