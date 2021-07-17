package graphql

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"strconv"

	"github.com/go-playground/validator"
	e "github.com/no-de-lab/nodelab-server/error"
	gqlschema "github.com/no-de-lab/nodelab-server/graphql/generated"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	sm "github.com/no-de-lab/nodelab-server/internal/study/model"
	"gopkg.in/jeevatkm/go-model.v1"
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

// CreateStudy create a study
func (sr *StudyResolver) CreateStudy(ctx context.Context, email string, input gqlschema.CreateStudyInput) (*gqlschema.Study, error) {

	user, err := sr.StudyService.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(user.ID)

	if err != nil {
		return nil, err
	}

	si := &sm.CreateStudy{
		Name:         input.Name,
		Limit:        input.Limit,
		StartDate:    input.StartDate,
		FinishDate:   input.FinishDate,
		Summary:      input.Summary,
		Title:        input.Title,
		Content:      input.Content,
		ThumbnailURL: input.ThumbnailURL,
		LeaderID:     userID,
		Notice:       input.Notice,
	}
	study, err := sr.StudyService.CreateStudy(ctx, si)

	if err != nil {
		return nil, err
	}
	var gqlStudy gqlschema.Study
	model.Copy(&gqlStudy, study)

	return &gqlStudy, nil
}

// UpdateStudy update a study
func (sr *StudyResolver) UpdateStudy(ctx context.Context, email string, id int, input gqlschema.UpdateStudyInput) (*gqlschema.Study, error) {

	study, err := sr.StudyService.FindByID(ctx, id)
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return nil, err
	}

	if err != nil {
		return nil, e.NewInternalError("can not find study", err, http.StatusInternalServerError)
	}

	user, err := sr.StudyService.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(user.ID)

	if err != nil {
		return nil, err
	}

	if study.LeaderID != userID {
		return nil, e.NewBusinessError("no authorization", nil, 403)
	}

	su := &sm.UpdateStudy{
		ID:           id,
		Name:         input.Name,
		Limit:        input.Limit,
		StartDate:    input.StartDate,
		FinishDate:   input.FinishDate,
		Summary:      input.Summary,
		Title:        input.Title,
		Content:      input.Content,
		ThumbnailURL: input.ThumbnailURL,
		Status:       input.Status,
	}

	study, err = sr.StudyService.UpdateStudy(ctx, su)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	var gqlStudy gqlschema.Study
	model.Copy(&gqlStudy, study)

	return &gqlStudy, nil
}

// DeleteStudy delete study
func (sr *StudyResolver) DeleteStudy(ctx context.Context, email string, id int) (bool, error) {
	study, err := sr.StudyService.FindByID(ctx, id)
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return false, err
	}

	if err != nil {
		return false, e.NewInternalError("can not find study", err, http.StatusInternalServerError)
	}

	user, err := sr.StudyService.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	userID, err := strconv.Atoi(user.ID)

	if err != nil {
		return false, err
	}

	if study.LeaderID != userID {
		return false, e.NewBusinessError("no authorization", nil, 403)
	}

	flag, err := sr.StudyService.DeleteStudy(ctx, id)
	if err != nil {
		return false, err
	}

	return flag, err
}
