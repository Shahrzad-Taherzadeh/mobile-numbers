package onmemory

import (
	"errors"
	"sync"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
)

var (
	Users []model.User
	mu    sync.RWMutex
)

func LoadInitUsers() {
	Create(model.User{
		Name:       "Shahrzad",
		FamilyName: "Taherzadeh",
		Age:        17,
		IsMarried:  false,
	})
}

func GetAll() []model.User {
	mu.RLock()
	defer mu.RUnlock()

	cpy := make([]model.User, len(Users))
	copy(cpy, Users)
	return cpy
}

func GetByID(id int) (model.User, error) {
	mu.RLock()
	defer mu.RUnlock()

	for _, u := range Users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func Create(u model.User) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	newID := len(Users) + 1
	u.ID = newID
	if u.MobileNumbers == nil {
		u.MobileNumbers = []model.MobileNumber{}
	}
	Users = append(Users, u)
	return newID, nil
}

func UpdateByID(id int, updated model.User) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range Users {
		if Users[i].ID == id {
			updated.ID = id
			if updated.MobileNumbers == nil {
				updated.MobileNumbers = Users[i].MobileNumbers
			}
			Users[i] = updated
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteByID(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range Users {
		if Users[i].ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func AddMobileNumber(id int, num model.MobileNumber) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range Users {
		if Users[i].ID == id {
			Users[i].MobileNumbers = append(Users[i].MobileNumbers, num)
			return nil
		}
	}
	return errors.New("user not found")
}

func DeleteMobileNumber(id int, number string) error {
	mu.Lock()
	defer mu.Unlock()

	for i := range Users {
		if Users[i].ID == id {
			for j, n := range Users[i].MobileNumbers {
				if n.Number == number {
					Users[i].MobileNumbers = append(Users[i].MobileNumbers[:j], Users[i].MobileNumbers[j+1:]...)
					return nil
				}
			}
			return errors.New("mobile number not found")
		}
	}
	return errors.New("user not found")
}
