package user

import "os/user"

func CreateUser(name string) *user.User {
	return &user.User{
		Name: name,
	}
}
