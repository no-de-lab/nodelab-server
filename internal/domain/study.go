package domain

import (
	"context"
	"time"
)

// Study is a struct to represent a nodelab study
type Study struct {
	ID         string    `db:"id"`
	CreatedAt  string    `db:"created_at"`
	UpdatedAt  string    `db:"updated_at"`
	StudyName  string    `db:"studyname"`
	Limit      int       `db:"limit"`
	StartDate  time.Time `db:"start_date"`
	FinishDate time.Time `db:"finish_date"`
	Summary    string    `db:"summary"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	LeaderID   int       `db:"leader_id"`
	Notice     string    `db:"notice"`
	Thumbnail  string    `db:"thumbnail"`
	Status     string    `db:"status"`
}

// StudyRepository is the respository layer for study
type StudyRepository interface {
	FindByID(context context.Context, id int) (study *Study, err error)
	FindByTitle(context context.Context, title string) (study *[]Study, err error)
}

// StudyService is the service layer for study
type StudyService interface {
	FindByID(context context.Context, id int) (study *Study, err error)
	FindByTitle(context context.Context, title string) (study *[]Study, err error)
}
