package domain

import (
	"context"

	am "github.com/no-de-lab/nodelab-server/internal/auth/model"
	"gopkg.in/guregu/null.v4"
)

// UserAccount is a struct for user account
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

// AuthRepository is the repository interface for any db operations for Auth
type AuthRepository interface {
	FindAccountByEmail(ctx context.Context, email string) (*UserAccount, error)
	CreateUserByEmail(ctx context.Context, user *UserAccount) error
	CreateUserBySocial(ctx context.Context, user *UserAccount) error
}

// AuthService is the service interface for Auth
type AuthService interface {
	Login(ctx context.Context, form *am.LoginModel) error
	SignupEmail(ctx context.Context, user *am.SignupEmailModel) (string, error)
	// SignupSocial(ctx context.Context, user *am.SignupSocialModel) (string, error)
	LoginSocial(ctx context.Context, form *am.LoginSocialModel) (string, *string, error)
	LoginEmail(ctx context.Context, email, password string) (string, error)
}
