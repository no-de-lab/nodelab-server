package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	um "github.com/no-de-lab/nodelab-server/internal/user/model"
)

type userDBRepository struct {
	DB *sqlx.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userDBRepository{
		db,
	}
}

// FindByID finds a user by ID
func (r *userDBRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	u := domain.User{}
	err := r.DB.GetContext(ctx, &u, FindByIDQuery, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// FindByEmail finds a user by email
func (r *userDBRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := domain.User{}
	err := r.DB.GetContext(ctx, &u, FindByEmailQuery, email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UpdateUser updates a user
func (r *userDBRepository) UpdateUser(ctx context.Context, userInfo *um.UserInfo) (*domain.User, error) {
	_, err := r.DB.NamedExecContext(ctx, updateUserByEmailQuery, userInfo)
	if err != nil {
		return nil, err
	}

	return r.FindByID(ctx, userInfo.ID)

}

// DeleteUser deletes a user
func (r *userDBRepository) DeleteUser(ctx context.Context, email string) error {
	_, err := r.DB.ExecContext(ctx, deleteUserByEmailQuery, email)
	if err != nil {
		return err
	}

	return nil
}
