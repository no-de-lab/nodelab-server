package model

import (
	"time"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
)

// UpdateStudy is used as a model for create a study
type UpdateStudy struct {
	ID           int                    `json:"id" db:"id"`
	Name         *string                `json:"name" db:"name"`
	Limit        *int                   `json:"limit" db:"limit"`
	StartDate    *time.Time             `json:"start_date" db:"start_date"`
	FinishDate   *time.Time             `json:"finish_date" db:"finish_date"`
	Summary      *string                `json:"summary" db:"summary"`
	Title        *string                `json:"title" db:"title"`
	Content      *string                `json:"content" db:"content"`
	Notice       *string                `json:"notice" db:"notice"`
	ThumbnailURL *string                `json:"thumbnail_url" db:"thumbnail_url"`
	Status       *gqlschema.StudyStatus `json:"status" db:"status"`
}
