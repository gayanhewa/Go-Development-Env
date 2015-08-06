package routes

import (
  "github.com/go-martini/martini"
  "fmt"
)

type Routes struct {}

func (r Routes) Load(&m **martini.ClassicMartini) {

  fmt.Println("Initializing the router")

  fmt.Println(m);

}

