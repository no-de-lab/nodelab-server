package graphql

import (
	"context"
	"strconv"

	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	"gopkg.in/jeevatkm/go-model.v1"
)

// UserResolver user resolver for graphql
type UserResolver struct {
	UserService domain.UserService
}

// NewUserResolver return new user resolver instance
func NewUserResolver(userService domain.UserService) *UserResolver {
	return &UserResolver{
		UserService: userService,
	}
}

// User find user by input id and returns
func (ur *UserResolver) User(ctx context.Context, _id string) (*gqlschema.User, error) {

	id, err := strconv.Atoi(_id)

	if err != nil {
		return nil, err
	}

	user, err := ur.UserService.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	var gqlUser gqlschema.User
	model.Copy(&gqlUser, user)

	// set mismatch type
	gqlUser.Intro = &user.Intro.String

	return &gqlUser, nil
}
