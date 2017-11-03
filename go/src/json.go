package main

import (
  "fmt"
  "encoding/json"
)

type Result struct {
  Err int `json:"err"`
  Msg string `json:"msg,omitempty"`
  Data interface {} `json:"data,omitempty"`
}

type ListData struct {
  Total int `json:"total"`
  Page int `json:"page"`
  Size int `json:"size"`
  List []interface {} `json:"list"`
}

type User struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

func main() {
  var _ string = `{
    "err": 0,
    "msg": "success",
    "data": {
      "total": 100,
      "page": 1,
      "size": 2,
      "list": [
        {
          "id": 1,
          "name": "one"
        },
        {
          "id": 2,
          "name": "two"
        }
      ]
    }
  }`

  var result map[string] interface {} = map[string] interface {} {
    "err": 0,
    "msg": "success",
    "data": map[string] interface {} {
      "total": 100,
      "page": 1,
      "size": 2,
      "list": []interface {} {
        map[string] interface {} {
          "id": 1,
          "name": "one",
        },
        map[string] interface {} {
          "id": 2,
          "name": "two",
        },
      },
    },
  }
  fmt.Println(result)
  json_data, _ := json.Marshal(result)
  json_string := string(json_data)
  fmt.Println(json_string)

  var result_data Result
  json.Unmarshal(json_data, &result_data)
  fmt.Println(result_data)
  json_data2, _ := json.Marshal(result_data)
  json_string2 := string(json_data2)
  fmt.Println(json_string2)

  var result2 Result = Result {
    Err: 0,
    Msg: "success",
    Data: ListData {
      Total: 100,
      Page: 1,
      Size: 2,
      List: []interface {} {
        User { Id: 1, Name: "one" },
        User { Id: 2, Name: "two" },
      },
    },
  }
  json_data3, _ := json.Marshal(result2)
  json_string3 := string(json_data3)
  fmt.Println(json_string3)
}
