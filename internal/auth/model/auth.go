package model

import "gopkg.in/guregu/null.v4"

type LoginModel struct {
	Email    string
	Password string
}

// SignupEmailModel is used to create users via email
type SignupEmailModel struct {
	Email    string      `json:"email" validate:"required,email"`
	Password null.String `json:"password" validate:"gte=6"`
}

// SignupSocialModel is used to create users via social account
type SignupSocialModel struct {
	Email string `json:"email" validate:"required,email"`
}
