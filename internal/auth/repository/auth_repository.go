package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/no-de-lab/nodelab-server/internal/domain"
)

type authRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) domain.AuthRepository {
	return &authRepository{
		db,
	}
}

func (r *authRepository) FindAccountByEmail(ctx context.Context, email string) (*domain.UserAccount, error) {
	u := domain.UserAccount{}
	err := r.DB.GetContext(ctx, &u, findAccountByEmailQuery, email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return &u, err
	}

	return &u, nil
}

// CreateUserByEmail creates user by email
func (r *authRepository) CreateUserByEmail(context context.Context, user *domain.UserAccount) error {
	_, err := r.DB.NamedExecContext(context, createUserByEmailQuery, user)
	if err != nil {
		return err
	}

	return nil
}

// CreateUserBySocial creates user by social account
func (r *authRepository) CreateUserBySocial(context context.Context, user *domain.UserAccount) error {
	_, err := r.DB.NamedExecContext(context, createUserBySocialQuery, user)
	if err != nil {
		return err
	}

	return nil
}
