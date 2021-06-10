package graphql

import (
	"context"

	"github.com/go-playground/validator"
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	"gopkg.in/jeevatkm/go-model.v1"
	"strconv"
)

// StudyResolver study resolver for graphql
type StudyResolver struct {
	Validator    validator.Validate
	StudyService domain.StudyService
}

// NewStudyResolver return new study resolver instance
func NewStudyResolver(validator validator.Validate, studyService domain.StudyService) *StudyResolver {
	return &StudyResolver{
		Validator:    validator,
		StudyService: studyService,
	}
}

// FindByID find study by ID
func (sr *StudyResolver) FindByID(ctx context.Context, _id string) (*gqlschema.Study, error) {

	id, err := strconv.Atoi(_id)

	if err != nil {
		return nil, err
	}

	study, err := sr.StudyService.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	var gqlStudy gqlschema.Study
	model.Copy(&gqlStudy, study)

	return &gqlStudy, nil
}

// FindByTitle find study by Title
func (sr *StudyResolver) FindByTitle(ctx context.Context, title string) ([]*gqlschema.Study, error) {

	study, err := sr.StudyService.FindByTitle(ctx, title)

	if err != nil {
		return nil, err
	}

	var gqlStudy []*gqlschema.Study
	model.Copy(&gqlStudy, study)

	return gqlStudy, nil
}
