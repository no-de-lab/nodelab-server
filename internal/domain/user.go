package domain

import (
	"context"
	"github.com/no-de-lab/nodelab-server/internal/user/model"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID             int64       `db:"id"`
	Email          string      `db:"email"`
	Username       string      `db:"username"`
	Password       null.String `db:"password"`
	Intro          null.String `db:"intro"`
	ProfileImageID null.Int    `db:"profile_image_id"`
	CreatedAt      string      `db:"created_at"`
	UpdatedAt      string      `db:"updated_at"`
}

type UserRepository interface {
	FindById(context context.Context, id int) (user *User, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
	CreateUser(context context.Context, user *User) (err error)
}

type UserService interface {
	FindById(context context.Context, id int) (user *User, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
	CreateUser(context context.Context, user *model.CreateUserModel) error
}
