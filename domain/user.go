package domain

import "context"

type User struct {
	Email        string `db:"email"`
	Username     string `db:"username"`
	Password     string `db:"password"`
	Intro        string `db:"intro"`
	ProfileImage int    `db:"profile_image"`
}

type UserRepository interface {
	CreateUser(context context.Context, user *User) (err error)
}

type UserService interface {
	CreateUser(context context.Context, user *User) error
}
