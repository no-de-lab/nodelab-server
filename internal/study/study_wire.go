package study

import (
	"github.com/google/wire"
	"github.com/no-de-lab/nodelab-server/internal/study/delivery/graphql"
	"github.com/no-de-lab/nodelab-server/internal/study/delivery/http"
	"github.com/no-de-lab/nodelab-server/internal/study/repository"
	"github.com/no-de-lab/nodelab-server/internal/study/service"
)

// StudySet a study domain instance set
var StudySet = wire.NewSet(repository.NewStudyRepository, service.NewStudyService, http.NewStudyHandler, graphql.NewStudyResolver)
