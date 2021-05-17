package model

// UserInfo is used to update user info
type UserInfo struct {
	ID              int64   `json:"id"`
	Email           string  `json:"email" db:"email"`
	Username        *string `json:"username" db:"username"`
	Intro           *string `json:"intro" db:"intro"`
	Position        *string `json:"position" db:"position"`
	Interest        *string `json:"interest" db:"interest"`
	ProfileImageUrl *string `json:"profile_image_url" db:"profile_image_url"`
	GithubURL       *string `json:"github_url" db:"github_url"`
	CreatedAt       string  `json:"createdAt" db:"created_at"`
	UpdatedAt       string  `json:"updatedAt" db:"updated_at"`
}
