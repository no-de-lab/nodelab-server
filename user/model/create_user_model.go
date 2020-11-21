package model

import "gopkg.in/guregu/null.v4"

type CreateUserModel struct {
	Email    null.String `json:"email"`
	Username null.String `json:"username"`
	Password null.String `json:"password"`
	Intro    null.String `json:"intro"`
}
