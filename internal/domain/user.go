package domain

import (
	"context"

	"github.com/no-de-lab/nodelab-server/internal/user/model"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID             string      `db:"id"`
	Email          string      `db:"email"`
	Username       null.String `db:"username"`
	Intro          null.String `db:"intro"`
	GithubURL      null.String `db:"github_url"`
	Position       null.String `db:"position"`
	Interest       null.String `db:"interest"`
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
