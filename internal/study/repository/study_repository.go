package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/no-de-lab/nodelab-server/internal/domain"
)

type studyDBRepository struct {
	DB *sqlx.DB
}

// NewStudyRepository creates a new study repository
func NewStudyRepository(db *sqlx.DB) domain.StudyRepository {
	return &studyDBRepository{
		db,
	}
}

// FindByID finds a study by ID
func (r *studyDBRepository) FindByID(ctx context.Context, id int) (study *domain.Study, err error) {
	s := domain.Study{}
	err = r.DB.GetContext(ctx, &s, FindByIDQuery, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *studyDBRepository) FindByTitle(ctx context.Context, title string) (study *[]domain.Study, err error) {
	s := []domain.Study{}
	err = r.DB.GetContext(ctx, &s, FindByTitleQuery, title)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}
