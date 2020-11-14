package dto

import "gopkg.in/guregu/null.v4"

type CreateUserDto struct {
	Email    null.String `json:"email"`
	Username null.String `json:"username"`
	Password null.String `json:"password"`
	Intro    null.String `json:"intro"`
}
