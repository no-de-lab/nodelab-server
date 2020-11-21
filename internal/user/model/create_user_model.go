package model

import "gopkg.in/guregu/null.v4"

type CreateUserModel struct {
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Password null.String `json:"password"`
	Intro    null.String `json:"intro"`
}
