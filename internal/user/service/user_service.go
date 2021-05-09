package service

import (
	"context"
	"time"

	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/internal/domain"
)

type userService struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

// NewUserService returns a new UserService instance
func NewUserService(userRepository domain.UserRepository, config *config.Configuration) domain.UserService {
	return &userService{
		userRepository,
		time.Duration(config.Context.Timeout) * time.Second,
	}
}

// FindByID finds user by id
func (s *userService) FindByID(c context.Context, id int) (user *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err = s.userRepository.FindByID(ctx, id)

	if err != nil {
		return
	}

	return
}

// FindByEmail finds user by email
func (s *userService) FindByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
