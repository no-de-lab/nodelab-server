package http

import "github.com/no-de-lab/nodelab-server/internal/domain"

type StudyHandler struct {
	StudyService domain.StudyService
}

func NewStudyHandler(ss domain.StudyService) *StudyHandler {
	return &StudyHandler{
		StudyService: ss,
	}
}
