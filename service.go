package service

import (
	"time"
)

type User struct {
	ID       int64  `json:"id,omitempty"`
	Nickname string `json:"nickname"`
}
type UserService interface {
	Add(User) User
	find() []User
}

type userService struct {
	data map[int64]User
}

func (s *userService) Add(u User) User {
	u.ID = time.Now().Unix()
	s.data[u.ID] = u
	return u
}

func (s *userService) find() []User {
	list := []User{}
	for _, v := range s.data {
		list = append(list, v)
	}
	return list
}
func NewUserService() UserService {
	data := make(map[int64]User)
	return &userService{data}
}
