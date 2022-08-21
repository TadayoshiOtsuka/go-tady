package entity

type User struct {
	ID   int
	Name string
}

func NewUser(id int, name string) (*User, error) {
	return &User{
		ID:   id,
		Name: name,
	}, nil
}
