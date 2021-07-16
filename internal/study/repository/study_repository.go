package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	sm "github.com/no-de-lab/nodelab-server/internal/study/model"
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
func (r *studyDBRepository) FindByID(ctx context.Context, id int) (*domain.Study, error) {
	s := domain.Study{}
	err := r.DB.GetContext(ctx, &s, FindByIDQuery, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// !TODO : scan error on column 'limit' converting type []uint8 to int error
	// if err != nil {
	// 	return nil, err
	// }

	return &s, nil
}

// FindByTitle find studies by title
func (r *studyDBRepository) FindByTitle(ctx context.Context, title string) (*[]domain.Study, error) {
	s := []domain.Study{}
	err := r.DB.GetContext(ctx, &s, FindByTitleQuery, title)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

// CreateStudy create study
func (r *studyDBRepository) CreateStudy(ctx context.Context, input *sm.CreateStudy) (*domain.Study, error) {
	res, err := r.DB.NamedExecContext(ctx, CreateStudyQuery, input)
	if err != nil {
		return nil, err
	}

	studyID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.FindByID(ctx, int(studyID))
}

// UpdateStudy update study
func (r *studyDBRepository) UpdateStudy(ctx context.Context, input *sm.UpdateStudy) (*domain.Study, error) {
	_, err := r.DB.NamedExecContext(ctx, UpdateStudyQuery, input)
	if err != nil {
		return nil, err
	}

	return r.FindByID(ctx, input.ID)
}

// DeleteStudy delete study
func (r *studyDBRepository) DeleteStudy(ctx context.Context, id int) (flag bool, err error) {
	_, err = r.DB.ExecContext(ctx, DeleteStudyQuery, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// FindByEmail finds a user by email
func (r *studyDBRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
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
