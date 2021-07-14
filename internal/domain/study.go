package domain

import (
	"context"
	"time"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	sm "github.com/no-de-lab/nodelab-server/internal/study/model"
)

// Study is a struct to represent a nodelab study
type Study struct {
	ID         string                `json:"id" db:"id"`
	CreatedAt  time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at" db:"updated_at"`
	Name       string                `json:"name" db:"name"`
	Limit      int                   `json:"limit" db:"limit"`
	StartDate  time.Time             `json:"start_date" db:"start_date"`
	FinishDate time.Time             `json:"finish_date" db:"finish_date"`
	Summary    string                `json:"summary" db:"summary"`
	Title      string                `json:"title" db:"title"`
	Content    string                `json:"content" db:"content"`
	LeaderID   int                   `json:"leader_id" db:"leader_id"`
	Notice     string                `json:"notice" db:"notice"`
	Thumbnail  string                `json:"thumbnail" db:"thumbnail_url"`
	Status     gqlschema.StudyStatus `json:"status" db:"status"`
}

// StudyRepository is the respository layer for study
type StudyRepository interface {
	FindByID(context context.Context, id int) (study *Study, err error)
	FindByTitle(context context.Context, title string) (study *[]Study, err error)
	CreateStudy(context context.Context, input *sm.CreateStudy) (study *Study, err error)
	UpdateStudy(context context.Context, input *sm.UpdateStudy) (study *Study, err error)
	DeleteStudy(context context.Context, id int) (flag bool, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
}

// StudyService is the service layer for study
type StudyService interface {
	FindByID(context context.Context, id int) (study *Study, err error)
	FindByTitle(context context.Context, title string) (study *[]Study, err error)
	CreateStudy(context context.Context, input *sm.CreateStudy) (study *Study, err error)
	UpdateStudy(context context.Context, input *sm.UpdateStudy) (study *Study, err error)
	DeleteStudy(context context.Context, id int) (flag bool, err error)
	FindByEmail(context context.Context, email string) (user *User, err error)
}
