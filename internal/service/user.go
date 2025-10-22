package service

import (
	"errors"
	"sync"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
	onmemory "github.com/Golang-Training-entry-3/mobile-numbers/internal/repository/on-memory"
)

var mu sync.Mutex

func GetUserList() ([]model.User, error) {
	mu.Lock()
	defer mu.Unlock()
	return onmemory.Users, nil
}

func GetUserByID(id int) (model.User, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range onmemory.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func CreateUser(user model.User) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	newRandomId := len(onmemory.Users) + 1
	user.ID = newRandomId
	onmemory.Users = append(onmemory.Users, user)
	return newRandomId, nil
}

func UpdateUserByID(id int, updated model.User) error {
	mu.Lock()
	defer mu.Unlock()

	for i, user := range onmemory.Users {
		if user.ID == id {
			updated.ID = id
			onmemory.Users[i] = updated
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteUserByID(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, user := range onmemory.Users {
		if user.ID == id {
			onmemory.Users = append(onmemory.Users[:i], onmemory.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func AddMobileNumber(id int, number model.MobileNumber) error {
	mu.Lock()
	defer mu.Unlock()

	for i, user := range onmemory.Users {
		if user.ID == id {
			onmemory.Users[i].MobileNumbers = append(onmemory.Users[i].MobileNumbers, number)
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteMobileNumber(id int, number string) error {
	mu.Lock()
	defer mu.Unlock()

	for i, user := range onmemory.Users {
		if user.ID == id {
			for j, num := range user.MobileNumbers {
				if num.Number == number {
					onmemory.Users[i].MobileNumbers = append(user.MobileNumbers[:j], user.MobileNumbers[j+1:]...)
					return nil
				}
			}
			return errors.New("mobile number not found")
		}
	}
	return errors.New("user not found")
}
