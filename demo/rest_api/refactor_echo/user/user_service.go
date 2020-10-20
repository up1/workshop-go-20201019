package user

import "context"

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

type UserServicer interface {
	GetUser(c context.Context) (User, error)
}

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

func (us UserService) GetUser(c context.Context) (User, error) {
	user := User{}
	return user, nil
}
