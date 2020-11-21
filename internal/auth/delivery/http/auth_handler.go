package http

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	e "github.com/no-de-lab/nodelab-server/error"
	"github.com/no-de-lab/nodelab-server/internal/domain"
	userError "github.com/no-de-lab/nodelab-server/internal/user/error"
	"github.com/no-de-lab/nodelab-server/internal/user/model"
	"net/http"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(service domain.AuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}

func (a *AuthHandler) SetupRoutes(mainRouter *mux.Router) {
	r := mainRouter.PathPrefix("/auth").Subrouter()
	r.HandleFunc("/signup", a.Signup).Methods("POST")
}

func (a *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user model.CreateUserModel

	context := r.Context()
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, e.ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	err = a.authService.Signup(context, &user)

	if err != nil {
		if errors.Is(err, userError.ErrUserAlreadyExists) {
			http.Error(w, userError.ErrUserAlreadyExists.Error(), http.StatusConflict)
			return
		}
		http.Error(w, e.ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
