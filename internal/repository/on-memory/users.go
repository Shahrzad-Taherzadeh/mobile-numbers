package onmemory

import (
	"errors"
	"sync"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
)

var (
	users []model.User
	mu    sync.RWMutex
)

func LoadInitUsers() {
	Create(model.User{
		Name:       "Ashva",
		FamilyName: "Patel",
		Age:        24,
		IsMarried:  false,
	})
}

func GetAll() []model.User {
	mu.RLock()
	defer mu.RUnlock()

	cpy := make([]model.User, len(users))
	copy(cpy, users)
	return cpy
}

func GetByID(id int) (model.User, error) {
	mu.RLock()
	defer mu.RUnlock()

	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func Create(u model.User) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	newID := len(users) + 1
	u.ID = newID
	if u.MobileNumbers == nil {
		u.MobileNumbers = []model.MobileNumber{}
	}
	users = append(users, u)
	return newID, nil
}

func UpdateByID(id int, updated model.User) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range users {
		if users[i].ID == id {
			updated.ID = id
			if updated.MobileNumbers == nil {
				updated.MobileNumbers = users[i].MobileNumbers
			}
			users[i] = updated
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteByID(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func AddMobileNumber(id int, num model.MobileNumber) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range users {
		if users[i].ID == id {
			users[i].MobileNumbers = append(users[i].MobileNumbers, num)
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteMobileNumber(id int, number string) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range users {
		if users[i].ID == id {
			for j, n := range users[i].MobileNumbers {
				if n.Number == number {
					users[i].MobileNumbers = append(users[i].MobileNumbers[:j], users[i].MobileNumbers[j+1:]...)
					return nil
				}
			}
			return errors.New("mobile number not found")
		}
	}
	return errors.New("user not found")
}
