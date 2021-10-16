package chainofresponsibility

import "fmt"

type Chain interface {
	SetNextChain(nextChain Chain)
	Check(request Request) bool
}

// Request is the command object
type Request struct {
	Username string
	Password string
}

// UsernameCheck is the processing objects
type UsernameCheck struct {
	nextChain Chain
}

func (u *UsernameCheck) SetNextChain(nextChain Chain) {
	u.nextChain = nextChain
}

func (u *UsernameCheck) Check(request Request) bool {
	// Check request.Username
	if true {
		u.nextChain.Check(request)
	}
	return false
}

// PasswordCheck is the processing objects
type PasswordCheck struct {
	nextChain Chain
}

func (p *PasswordCheck) SetNextChain(nextChain Chain) {
	p.nextChain = nextChain
}

func (p *PasswordCheck) Check(request Request) bool {
	// Check request.Password
	if true {
		p.nextChain.Check(request)
	}
	return false
}

func Run() {
	usrCheck := UsernameCheck{}
	pwdCheck := PasswordCheck{}

	usrCheck.SetNextChain(&pwdCheck)
	valid := usrCheck.Check(Request{})
	if valid {
		fmt.Println("request valid")
	}
}
