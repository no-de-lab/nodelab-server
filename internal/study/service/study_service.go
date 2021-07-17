package service

import (
	"context"
	"time"

	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	sm "github.com/no-de-lab/nodelab-server/internal/study/model"
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
		return nil, err
	}

	return study, err
}

// FindByTitle finds user by id
func (s *studyService) FindByTitle(c context.Context, name string) (study *[]domain.Study, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	study, err = s.studyRepository.FindByTitle(ctx, name)

	if err != nil {
		return nil, err
	}

	return study, err
}

// CreateStudy create study
func (s *studyService) CreateStudy(ctx context.Context, input *sm.CreateStudy) (*domain.Study, error) {
	study, err := s.studyRepository.CreateStudy(ctx, input)
	if err != nil {
		return nil, err
	}

	return study, err
}

// UpdateStudy update study
func (s *studyService) UpdateStudy(ctx context.Context, input *sm.UpdateStudy) (*domain.Study, error) {
	study, err := s.studyRepository.UpdateStudy(ctx, input)
	if err != nil {
		return nil, err
	}

	return study, err
}

// DeleteStudy delete study
func (s *studyService) DeleteStudy(ctx context.Context, id int) (flag bool, err error) {
	flag, err = s.studyRepository.DeleteStudy(ctx, id)
	if err != nil {
		return flag, err
	}

	return flag, nil
}

// FindByEmail finds user by email
func (s *studyService) FindByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.studyRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
