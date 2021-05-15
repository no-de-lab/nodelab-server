package model

import (
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"gopkg.in/guregu/null.v4"
)

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
type LoginSocialModel struct {
	Email       string             `json:"email" validate:"required,email"`
	AccessToken string             `json:"access_token" validate:"required,access_token"`
	Provider    gqlschema.Provider `json:"provider" validate:"required,provider"`
}
