package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidLogin = errors.New("invalid login")
	ErrUsernameTaken = errors.New("username taken")
)

type User struct {
	id int64
	//key string
}

func NewUser(username string, hash []byte) (*User, error) {
	exists, err := client.HExists(context.TODO(), "user:by-username", username).Result()
	if exists {
		return nil, ErrUsernameTaken
	}
	id, err := client.Incr(context.TODO(), "user:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(context.TODO(), key, "id", id)
	pipe.HSet(context.TODO(), key,"username", username)
	pipe.HSet(context.TODO(), key,"hash", hash)
	pipe.HSet(context.TODO(), "user:by-username", username, id)
	_, err = pipe.Exec(context.TODO())
	if err != nil {
		return nil, err
	}
	//return &User{key}, nil
	return &User{id}, nil
}

func (user *User) GetId() (int64, error) {
	//key := fmt.Sprintf("user:%d", user.id)
	//return client.HGet(context.TODO(), key, "id").Int64()
	return user.id, nil
}

func (user *User) GetUsername() (string, error) {
	key := fmt.Sprintf("user:%d", user.id)
	//return client.HGet(context.TODO(), user.key, "username").Result()
	return client.HGet(context.TODO(), key, "username").Result()
}

func (user *User) GetHash() ([]byte, error) {
	key := fmt.Sprintf("user:%d", user.id)
	//return client.HGet(context.TODO(), user.key, "hash").Bytes()
	return client.HGet(context.TODO(), key, "hash").Bytes()
}

func (user *User) Authenticate(password string) error {
	hash, err  := user.GetHash()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrInvalidLogin
	}
	return err
}

func GetUserById(id int64) (*User, error) {
	//key := fmt.Sprintf("user:%d", id)
	//return &User{key}, nil
	return &User{id}, nil
}

func GetUserByUsername(username string) (*User, error) {
	id, err := client.HGet(context.TODO(), "user:by-username", username).Int64()
	if err == redis.Nil {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	//key := fmt.Sprintf("user:%d", id)
	//return &User{key}, nil
	return GetUserById(id)
}

func AuthenticateUser(username string, password string) (*User, error) {
	//hash, err := client.Get(context.TODO(), "user:" + username).Bytes()
	//if err == redis.Nil {
	//	return ErrUserNotFound
	//} else if err != nil {
	//	return err
	//}
	//err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	//if err != nil {
	//	return ErrInvalidLogin
	//}
	//return nil
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, user.Authenticate(password)
}

func RegisterUser(username string, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return  err
	}
	_, err = NewUser(username, hash)
	//return client.Set(context.TODO(), "user:" + username, hash, 0).Err()
	return err
}