package testify

type UserModel struct {
	ID   int
	Name string
}

type UserInterface interface {
	Get(id int) (*UserModel, error)
}

type User struct {
	Dt UserInterface
}

func (u *User) UserName(id int) (string, error) {
	usr, err := u.Dt.Get(id)
	if usr == nil || err != nil {
		return "", err
	}
	return usr.Name, nil
}

func main() {}
