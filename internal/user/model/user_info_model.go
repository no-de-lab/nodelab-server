package model

import "gopkg.in/guregu/null.v4"

type UserInfoModel struct {
	ID              int64       `json:"id"`
	Email           string      `json:"email"`
	Username        string      `json:"username"`
	Intro           null.String `json:"intro"`
	ProfileImageUrl null.String `json:"profileImageUrl"`
	CreatedAt       string      `json:"createdAt"`
	UpdatedAt       string      `json:"updatedAt"`
}
