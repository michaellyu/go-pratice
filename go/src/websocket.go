package main

import (
  "github.com/gorilla/websocket"
  "net/http"
  "html/template"
  "log"
)

var (
  addr string = ":8080"
  upgrader = websocket.Upgrader {}
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
  c, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println("upgrade: ", err)
    return
  }
  defer c.Close()

  for {
    mt, message, err := c.ReadMessage()
    if err != nil {
      log.Println("read: ", err)
      break
    }
    log.Printf("Receive from clients: %s", message)
    err = c.WriteMessage(mt,  message)
    log.Printf("Send to clients: %s", message)
    if err != nil {
      log.Println("write: ", err)
      return
    }
  }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("websocket.html")
  t.Execute(w, "ws://" + r.Host + "/echo")
}

func main() {
  log.SetFlags(0)
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/echo", echoHandler)
  log.Fatal(http.ListenAndServe(addr, nil))
}
