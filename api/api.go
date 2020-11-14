package api

import "github.com/gorilla/mux"

type ApiHandler interface {
	SetupRoutes(r *mux.Router)
}
