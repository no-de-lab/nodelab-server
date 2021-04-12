package domain

import (
	"context"

	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"gopkg.in/guregu/null.v4"
)

type UserAccount struct {
	ID          int64       `db:"id"`
	Email       string      `db:"email"`
	Password    null.String `db:"password"`
	Provider    null.String `db:"provider"`
	ProviderID  null.String `db:"provider_id"`
	AccessToken null.String `db:"access_token"`
	CreatedAt   string      `db:"created_at"`
	UpdatedAt   string      `db:"updated_at"`
}

type AuthRepository interface {
	FindAccountByEmail(ctx context.Context, email string) (*UserAccount, error)
	CreateUserByEmail(ctx context.Context, user *UserAccount) error
	CreateUserBySocial(ctx context.Context, user *UserAccount) error
}
type AuthService interface {
	Login(ctx context.Context, form *am.LoginModel) error
	SignupEmail(ctx context.Context, user *am.SignupEmailModel) error
	SignupSocial(ctx context.Context, user *am.SignupSocialModel) error
	// TODO: add detail
	SocialLogin() error
}
