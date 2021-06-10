package service

import (
	"context"
	"time"

	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/internal/domain"
)

type studyService struct {
	studyRepository domain.StudyRepository
	timeout         time.Duration
}

// NewStudyService returns a new StudyService instance
func NewStudyService(studyRepository domain.StudyRepository, config *config.Configuration) domain.StudyService {
	return &studyService{
		studyRepository,
		time.Duration(config.Context.Timeout) * time.Second,
	}
}

// FindByID finds study by id
func (s *studyService) FindByID(c context.Context, id int) (study *domain.Study, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	study, err = s.studyRepository.FindByID(ctx, id)

	if err != nil {
		return
	}

	return study, err
}

// FindByTitle finds user by id
func (s *studyService) FindByTitle(c context.Context, name string) (study *[]domain.Study, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	study, err = s.studyRepository.FindByTitle(ctx, name)

	if err != nil {
		return
	}

	return study, err
}
