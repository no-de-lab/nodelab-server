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
func (r *userDBRepository) FindByID(ctx context.Context, id int) (user *domain.User, err error) {
	u := domain.User{}
	err = r.DB.GetContext(ctx, &u, FindByIDQuery, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return
	}

	return &u, nil
}

// FindByEmail finds a user by email
func (r *userDBRepository) FindByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	u := domain.User{}
	err = r.DB.GetContext(ctx, &u, FindByEmailQuery, email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return
	}

	return &u, nil
}

// UpdateUser updates a user
func (r *userDBRepository) UpdateUser(ctx context.Context, userInfo *um.UserInfo) error {
	_, err := r.DB.NamedExecContext(ctx, updateUserByEmailQuery, userInfo)
	if err != nil {
		return err
	}

	return nil
}
