package repository

import (
	"context"
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

func (r *userDBRepository) CreateUser(context context.Context, user *domain.User) (err error) {
	query := `INSERT INTO user(email, username, password, intro, profile_image) VALUES(:email, :username, :password, :intro, :profile_image)`
	_, err = r.Conn.NamedExecContext(context, query, user)

	if err != nil {
		return
	}

	return nil
}
