package bll

type User struct {
	Name string
}

func GetUsers() ([]User, error) {
	return []User{
		{"U1"},
		{"U2"},
	}, nil
}
