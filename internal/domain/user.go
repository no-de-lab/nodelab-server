package domain

import (
	"context"

	um "github.com/no-de-lab/nodelab-server/internal/user/model"
)

// User is a struct to represent a nodelab user
type User struct {
	ID             string  `db:"id"`
	Email          string  `db:"email"`
	Username       *string `db:"username"`
	Intro          *string `db:"intro"`
	GithubURL      *string `db:"github_url"`
	Position       *string `db:"position"`
	Interest       *string `db:"interest"`
	ProfileImageID *int    `db:"profile_image_id"`
	CreatedAt      string  `db:"created_at"`
	UpdatedAt      string  `db:"updated_at"`
}

// UserRepository is the repository layer for user
type UserRepository interface {
	FindByID(context context.Context, id int) (user *User, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
	UpdateUser(context context.Context, userInfo *um.UserInfo) (err error)
	DeleteUser(context context.Context, email string) (err error)
}

// UserService is the service layer for user
type UserService interface {
	FindByID(context context.Context, id int) (user *User, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
	UpdateUser(context context.Context, userInfo *um.UserInfo) (user *User, err error)
	DeleteUser(context context.Context, email string) (err error)
}
