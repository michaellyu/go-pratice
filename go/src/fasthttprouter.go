package main

import (
  "fmt"
  "log"
  "github.com/buaazp/fasthttprouter"
  "github.com/valyala/fasthttp"
)

func main() {
  router := fasthttprouter.New()
  router.GET("/", Index)
  router.GET("/hello/:name", Hello)
  router.GET("/params", Params)
  router.POST("/post", Post)

  log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

func Index(ctx *fasthttp.RequestCtx) {
  fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
  fmt.Fprintf(ctx, "Hello, %s\n", ctx.UserValue("name"))
}

func Params(ctx *fasthttp.RequestCtx) {
  var json string
  args := ctx.QueryArgs()
  json += "{\n"
  json += "  \"username\": \"" + string(args.Peek("username")[:]) + "\",\n"
  json += "  \"password\": \"" + string(args.Peek("password")[:]) + "\",\n"
  json += "  \"remember\": \"" + string(args.Peek("remember")[:]) + "\"\n"
  json += "}\n"
  ctx.WriteString(json)
}

func Post(ctx *fasthttp.RequestCtx) {
  var json string
  args := ctx.PostArgs()
  json += "{\n"
  json += "  \"username\": \"" + string(args.Peek("username")[:]) + "\",\n"
  json += "  \"password\": \"" + string(args.Peek("password")[:]) + "\",\n"
  json += "  \"remember\": \"" + string(args.Peek("remember")[:]) + "\"\n"
  json += "}\n"
  ctx.WriteString(json)
}
