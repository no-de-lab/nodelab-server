package model

import "gopkg.in/guregu/null.v4"

// CreateUser is used as a model for creating a user
type CreateUser struct {
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Password null.String `json:"password"`
	Intro    null.String `json:"intro"`
}
