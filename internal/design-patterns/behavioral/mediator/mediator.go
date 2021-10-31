package mediator

import (
	"fmt"
	"reflect"
)

// Button is the component
type Button struct {
	Name    string
	onClick func()
}

func (b *Button) GetName() string {
	return b.Name
}

func (b *Button) SetFunc(f func()) {
	b.onClick = f
}

// Input is the component
type Input struct {
	Value string
}

func (i *Input) GetValue() string {
	return i.Value
}

// LoginPage is the Mediator
type LoginPage struct {
	UsernameInput *Input
	PasswordInput *Input
	LoginButton   *Button
	RegButton     *Button
}

func (l *LoginPage) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, l.LoginButton):
		fmt.Println(l.UsernameInput.GetValue())
		fmt.Println(l.PasswordInput.GetValue())
		l.LoginButton.onClick()
	case reflect.DeepEqual(component, l.RegButton):
		fmt.Println(l.UsernameInput.GetValue())
		fmt.Println(l.PasswordInput.GetValue())
		l.LoginButton.onClick()
	}
}

func Run() {
	usernameInput := Input{}
	passwordInput := Input{}
	loginButton := Button{Name: "Login"}
	regButton := Button{Name: "Register"}
	l := &LoginPage{
		UsernameInput: &usernameInput,
		PasswordInput: &passwordInput,
		LoginButton:   &loginButton,
		RegButton:     &regButton,
	}

	l.HandleEvent(loginButton)
	l.HandleEvent(loginButton)
}
