package main

type User interface {
	getAge() int
}

type UserService struct {
	User User
}

func NewUserService(user User) *UserService {
	return &UserService{
		User: user,
	}
}

type user struct {
	ID   string
	Name string
	Age  int
}

func (u *user) getAge() int {
	return u.Age
}
