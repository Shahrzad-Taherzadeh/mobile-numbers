package service

import (
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/repository"
)

var repo repository.UserRepository

func SetRepository(userRepo repository.UserRepository) {
	repo = userRepo
}

func GetUserList() ([]model.User, error) {
	return repo.GetAll()
}

func GetUserByID(id int) (model.User, error) {
	return repo.GetByID(id)
}

func CreateUser(user model.User) (int, error) {
	return repo.Create(user)
}

func UpdateUserByID(id int, updated model.User) error {
	return repo.UpdateByID(id, updated)
}

func DeleteUserByID(id int) error {
	return repo.DeleteByID(id)
}

func AddMobileNumber(id int, number model.MobileNumber) error {
	return repo.AddMobileNumber(id, number)
}

func DeleteMobileNumber(id int, number string) error {
	return repo.DeleteMobileNumber(id, number)
}

