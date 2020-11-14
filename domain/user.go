package domain

import (
	"context"
	"github.com/no-de-lab/nodelab-server/user/dto"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id           null.Int    `db:"id"`
	Email        null.String `db:"email"`
	Username     null.String `db:"username"`
	Password     null.String `db:"password"`
	Intro        null.String `db:"intro"`
	ProfileImage null.Int    `db:"profile_image"`
	CreatedAt    null.String `db:"created_at"`
	UpdatedAt    null.String `db:"updated_at"`
}

type UserRepository interface {
	FindById(context context.Context, id int) (user *User, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
	CreateUser(context context.Context, user *User) (err error)
}

type UserService interface {
	FindById(context context.Context, id int) (user *User, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
	CreateUser(context context.Context, user *dto.CreateUserDto) error
}
