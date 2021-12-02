package decorator

// Use of the decorator can provide enhanced logic.

type IUser interface {
	Login(username, password string) error
}

type User struct{}

func (u *User) Login(username, password string) error {
	// Login
	return nil
}

type VIPUser struct {
	user *User
}

func (v *VIPUser) showBadge() {}

func (v *VIPUser) Login(username, password string) error {
	err := v.user.Login(username, password)
	if err != nil {
		return err
	}

	v.showBadge()

	return nil
}

func Run() {
	vipUser := VIPUser{
		user: &User{},
	}
	vipUser.Login("username", "password")
}
