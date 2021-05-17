package graphql

import (
	"context"
	"strconv"

	"github.com/go-playground/validator"
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	um "github.com/no-de-lab/nodelab-server/internal/user/model"
	log "github.com/sirupsen/logrus"
	"gopkg.in/jeevatkm/go-model.v1"
)

// UserResolver user resolver for graphql
type UserResolver struct {
	Validator   validator.Validate
	UserService domain.UserService
}

// NewUserResolver return new user resolver instance
func NewUserResolver(validator validator.Validate, userService domain.UserService) *UserResolver {
	return &UserResolver{
		Validator:   validator,
		UserService: userService,
	}
}

// User find user by input id and returns
func (ur *UserResolver) User(ctx context.Context, _id string) (*gqlschema.User, error) {

	id, err := strconv.Atoi(_id)

	if err != nil {
		return nil, err
	}

	user, err := ur.UserService.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	var gqlUser gqlschema.User
	model.Copy(&gqlUser, user)

	// set mismatch type
	gqlUser.Intro = user.Intro

	return &gqlUser, nil
}

// Me gets the current users information from the token payload
func (ur *UserResolver) Me(ctx context.Context, email string) (*gqlschema.User, error) {
	user, err := ur.UserService.FindByEmail(ctx, email)
	if err != nil {
		log.WithError(err).Errorf("failed to get user for email: %s", email)
		return nil, err
	}

	var gqlUser gqlschema.User
	model.Copy(&gqlUser, user)

	return &gqlUser, nil
}

// UpdateUser updates the current user's information with the given payload
func (ur *UserResolver) UpdateUser(ctx context.Context, email string, payload *gqlschema.UpdateUserInput) (*gqlschema.User, error) {
	uim := &um.UserInfo{
		Email:     email,
		Username:  payload.Username,
		Intro:     payload.Intro,
		Position:  payload.Position,
		Interest:  payload.Interest,
		GithubURL: payload.GithubURL,
	}

	user, err := ur.UserService.UpdateUser(ctx, uim)
	if err != nil {
		log.WithError(err).Errorf("failed to update user: %s", email)
		return nil, err
	}

	var gqlUser gqlschema.User
	model.Copy(&gqlUser, user)

	return &gqlUser, nil
}

// DeleteUser deletes the current user's information
func (ur *UserResolver) DeleteUser(ctx context.Context, email string) (string, error) {
	err := ur.UserService.DeleteUser(ctx, email)
	if err != nil {
		log.WithError(err).Errorf("failed to delete user: %s", email)
		return "", err
	}

	return email, nil
}
