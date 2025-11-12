package repository

import "github.com/Golang-Training-entry-3/mobile-numbers/internal/model"

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (model.User, error)
	Create(u model.User) (int, error)
	UpdateByID(id int, updated model.User) error
	DeleteByID(id int) error
	AddMobileNumber(id int, num model.MobileNumber) error
	DeleteMobileNumber(id int, number string) error
}