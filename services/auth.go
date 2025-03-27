package services

import "fmt"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{}

func RegisterUser(newUser User) error {
	for _, user := range users {
		if user.Username == newUser.Username {
			return fmt.Errorf("username already exists")
		}
	}
	users = append(users, newUser)
	return nil
}

func LoginUser(username, password string) (*User, error) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("invalid credentials")
}
