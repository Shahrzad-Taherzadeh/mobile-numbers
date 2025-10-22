package service

import (
	"errors"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
	onmemory "github.com/Golang-Training-entry-3/mobile-numbers/internal/repository/on-memory"
)

func GetUserList() ([]model.User, error) {
	users := onmemory.GetAll()
	return users, nil
}

func GetUserByID(id int) (model.User, error) {
	return onmemory.GetByID(id)
}

func CreateUser(user model.User) (int, error) {
	return onmemory.Create(user)
}

func UpdateUserByID(id int, updated model.User) error {
	return onmemory.UpdateByID(id, updated)
}

func DeleteUserByID(id int) error {
	return onmemory.DeleteByID(id)
}

func AddMobileNumber(id int, number model.MobileNumber) error {
	return onmemory.AddMobileNumber(id, number)
}

func DeleteMobileNumber(id int, number string) error {
	return onmemory.DeleteMobileNumber(id, number)
}
