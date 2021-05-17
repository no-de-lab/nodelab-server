package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/no-de-lab/nodelab-server/config"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	um "github.com/no-de-lab/nodelab-server/internal/user/model"
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

// UpdateUser updates the user with the given payload
func (s *userService) UpdateUser(c context.Context, userInfo *um.UserInfo) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.userRepository.UpdateUser(ctx, userInfo)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			return nil, fmt.Errorf("username: %s already exists", *userInfo.Username)
		}
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	user, err := s.userRepository.FindByEmail(ctx, userInfo.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to read updated user: %v", err)
	}

	return user, nil
}

// DeleteUser deletes the user for the given email
func (s *userService) DeleteUser(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.userRepository.DeleteUser(ctx, email)
	return err
}
