package controllers

import (
  "repo"
  "fmt"
  "encoding/json"
)

type UserController struct {

}

func (u *UserController) GetUsers() ([]byte,error) {
  var uRepo = repo.UserRepository{}

  var result =  uRepo.GetUsers()

  fmt.Println(result)


  res, err := json.Marshal(uRepo.GetUsers())

  if err != nil {

    fmt.Println(err)
    return nil,err
  }

  fmt.Println(res)
  return res,nil
}


