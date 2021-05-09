package service

import (
	"context"
	"time"

	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	um "github.com/no-de-lab/nodelab-server/internal/user/model"
	"gopkg.in/jeevatkm/go-model.v1"
)

type userService struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewUserService(userRepository domain.UserRepository, config *config.Configuration) domain.UserService {
	return &userService{
		userRepository,
		time.Duration(config.Context.Timeout) * time.Second,
	}
}

func (s *userService) FindById(c context.Context, id int) (user *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err = s.userRepository.FindById(ctx, id)

	if err != nil {
		return
	}

	return
}

func (s *userService) FindByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(c context.Context, user *um.CreateUserModel) (err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	var userModel domain.User
	errs := model.Copy(&userModel, user)

	if errs != nil {
		return errs[0]
	}

	err = s.userRepository.CreateUser(ctx, &userModel)
	if err != nil {
		return err
	}

	return nil
}
