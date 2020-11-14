package dto

import "gopkg.in/guregu/null.v4"

type UserInfoDto struct {
	Id              null.Int    `json:"id"`
	Email           null.String `json:"email"`
	Username        null.String `json:"username"`
	Intro           null.String `json:"intro"`
	ProfileImageUrl null.String `json:"profileImageUrl"`
	CreatedAt       null.String `json:"createdAt"`
	UpdatedAt       null.String `json:"updatedAt"`
}
