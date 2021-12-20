package facade

type UserService struct{}

func (u *UserService) Login() error {
	return nil
}

func (u *UserService) Register() error {
	return nil
}

func (u *UserService) LoginOrRegister() error {
	err := u.Login()
	if err != nil {
		return err
	}
	err = u.Login()
	if err != nil {
		return err
	}
	return nil
}
