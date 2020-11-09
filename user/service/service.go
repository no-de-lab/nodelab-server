package service

import (
	"context"
	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/domain"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewUserService(userRepository domain.UserRepository, config *config.Config) domain.UserService {
	return &userService{
		userRepository,
		time.Duration(config.Context.Timeout) * time.Second,
	}
}

func (s *userService) CreateUser(c context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err = s.userRepository.CreateUser(ctx, user)

	if err != nil {
		return
	}

	return nil
}
