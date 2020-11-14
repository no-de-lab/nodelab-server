package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/no-de-lab/nodelab-server/domain"
)

type userDBRepository struct {
	Conn *sqlx.DB
}

func NewUserRepository(Conn *sqlx.DB) domain.UserRepository {
	return &userDBRepository{
		Conn,
	}
}

func (r *userDBRepository) FindById(ctx context.Context, id int) (user *domain.User, err error) {
	u := domain.User{}
	err = r.Conn.GetContext(ctx, &u, FindByIdQuery, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return
	}

	return &u, nil
}

func (r *userDBRepository) FindByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	u := domain.User{}
	err = r.Conn.GetContext(ctx, &u, FindByEmailQuery, email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return
	}

	return &u, nil
}

func (r *userDBRepository) CreateUser(context context.Context, user *domain.User) (err error) {
	_, err = r.Conn.NamedExecContext(context, CreateUserQuery, user)

	if err != nil {
		return
	}

	return nil
}
