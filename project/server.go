package main

import (
  "github.com/go-martini/martini"
  "fmt"
  "controllers"
)

var (
   m *martini.Martini
   UserController *controllers.UserController
)

func main() {
  version := "1.0";
  fmt.Println("Starting version"+version)

  m := martini.Classic()

  m.Get("/users", UserController.GetUsers)

  m.Run()
}
