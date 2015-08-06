package repo

import (
  "models"
)


type UserInterface interface {
  GetUsers() []models.User
}

type UserRepository struct{}

func (u *UserRepository) GetUsers() []models.User{

  users_array := make([]models.User,2)
  users_array[0] = models.User{"Gayan"}
  users_array[1] = models.User{"Hewa"}

  return users_array
}
