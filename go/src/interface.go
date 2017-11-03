package main

import "fmt"

type Animal interface {
  run()
}

type Dog struct {
  name string
}

func(dog Dog) run() {
  fmt.Printf("%s is running\n", dog.name)
}

func animalRun(animal Animal) {
  animal.run()
}

func main() {
  var dog Dog = Dog { name: "Dog" }
  animalRun(dog)
}
