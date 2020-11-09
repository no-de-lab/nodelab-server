package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/no-de-lab/nodelab-server/domain"
	"net/http"
)

type UserHandler struct {
	UserService domain.UserService
}

func NewUserHandler(us domain.UserService) *mux.Router {
	handler := &UserHandler{
		UserService: us,
	}

	r := mux.NewRouter()
	r.HandleFunc("/user/create", handler.CreateUser).Methods("POST")

	return r
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	context := r.Context()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserService.CreateUser(context, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
