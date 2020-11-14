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
	query := `SELECT id, email, username, intro, profile_image, created_at, updated_at FROM user WHERE id = ?`

	err = r.Conn.GetContext(ctx, &u, query, id)

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
	query := `SELECT email, password FROM user WHERE email = ?`
	err = r.Conn.GetContext(ctx, &u, query, email)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return
	}

	return &u, nil
}

func (r *userDBRepository) CreateUser(context context.Context, user *domain.User) (err error) {
	query := `
		INSERT INTO user(email, username, password, intro, profile_image, created_at, updated_at) 
		VALUES(:email, :username, :password, :intro, :profile_image, now(), now())
	`
	_, err = r.Conn.NamedExecContext(context, query, user)

	if err != nil {
		return
	}

	return nil
}
