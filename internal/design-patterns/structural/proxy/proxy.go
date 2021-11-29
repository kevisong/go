package proxy

import (
	"fmt"
	"time"
)

type IUser interface {
	Login(username, password string) error
}

type User struct{}

func (u *User) Login(username, password string) error {
	// Login
	return nil
}

type UserProxy struct {
	user *User
}

func (u *UserProxy) Login(username, password string) error {
	// Before
	start := time.Now()

	err := u.user.Login(username, password)
	if err != nil {
		return err
	}

	// After
	end := time.Now().Sub(start)

	fmt.Printf("start: %v lasts: %v", start, end)
	return nil
}

func Run() {
	userProxy := UserProxy{
		user: &User{},
	}
	userProxy.Login("username", "password")
}
