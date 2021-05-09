package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/no-de-lab/nodelab-server/internal/user/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const timeFormat = "2006-01-02 15:04:05"

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "email", "username", "intro", "profile_image_id", "created_at", "updated_at"}).
		AddRow(1, "test@gmail.com", "test", "test intro", 0, time.Now().Format(timeFormat), time.Now().Format(timeFormat))

	mock.ExpectQuery("SELECT id").WillReturnRows(rows)

	repo := repository.NewUserRepository(sqlx.NewDb(db, "mysql"))

	user, err := repo.FindByID(context.TODO(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.ID, int64(1))
}
