package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type UserRepository struct {
	client *redis.Client
}

func NewUserRepository(client *redis.Client) *UserRepository {
	return &UserRepository{client: client}
}

func key(id int) string {
	return fmt.Sprintf("user:%d", id)
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	keys, err := r.client.Keys(ctx, "user:*").Result()
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return []model.User{}, nil
	}

	data, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0, len(data))
	for _, raw := range data {
		if raw == nil {
			continue
		}
		var u model.User
		if err := json.Unmarshal([]byte(raw.(string)), &u); err != nil {
			return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (model.User, error) {
	data, err := r.client.Get(ctx, key(id)).Result()
	if err == redis.Nil {
		return model.User{}, errors.New("user not found")
	}
	if err != nil {
		return model.User{}, err
	}

	var u model.User
	if err := json.Unmarshal([]byte(data), &u); err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *UserRepository) Create(u model.User) (int, error) {
	newID, err := r.client.Incr(ctx, "user:next_id").Result()
	if err != nil {
		return 0, err
	}
	u.ID = int(newID)
	
	if u.MobileNumbers == nil {
		u.MobileNumbers = []model.MobileNumber{}
	}

	serialized, err := json.Marshal(u)
	if err != nil {
		return 0, err
	}

	if err := r.client.Set(ctx, key(u.ID), serialized, 0).Err(); err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (r *UserRepository) UpdateByID(id int, updated model.User) error {
	current, err := r.GetByID(id)
	if err != nil {
		return err \
	}
	
	if updated.MobileNumbers == nil {
		updated.MobileNumbers = current.MobileNumbers
	}
	
	updated.ID = id

	serialized, err := json.Marshal(updated)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key(id), serialized, 0).Err()
}

func (r *UserRepository) DeleteByID(id int) error {
	count, err := r.client.Del(ctx, key(id)).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) AddMobileNumber(id int, num model.MobileNumber) error {
	_, err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		current, err := r.GetByID(id)
		if err != nil {
			return err
		}

		current.MobileNumbers = append(current.MobileNumbers, num)
		
		serialized, err := json.Marshal(current)
		if err != nil {
			return err
		}

		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key(id), serialized, 0)
			return nil
		})
		return err
	}, key(id))

	return err
}

func (r *UserRepository) DeleteMobileNumber(id int, number string) error {
	_, err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		current, err := r.GetByID(id)
		if err != nil {
			return err
		}

		found := false
		newNumbers := make([]model.MobileNumber, 0)
		for _, n := range current.MobileNumbers {
			if strings.EqualFold(n.Number, number) {
				found = true
				continue
			}
			newNumbers = append(newNumbers, n)
		}

		if !found {
			return errors.New("mobile number not found")
		}

		current.MobileNumbers = newNumbers
		
		serialized, err := json.Marshal(current)
		if err != nil {
			return err
		}

		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key(id), serialized, 0)
			return nil
		})
		return err
	}, key(id))

	return err
}