package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/no-de-lab/nodelab-server/domain"
	e "github.com/no-de-lab/nodelab-server/error"
	userError "github.com/no-de-lab/nodelab-server/user/error"
	um "github.com/no-de-lab/nodelab-server/user/model"
	"gopkg.in/jeevatkm/go-model.v1"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService domain.UserService
}

func NewUserHandler(us domain.UserService) *UserHandler {
	return &UserHandler{
		UserService: us,
	}
}

func (h *UserHandler) SetupRoutes(mainRouter *mux.Router) {
	r := mainRouter.PathPrefix("/users").Subrouter()
	r.HandleFunc("/{id}", h.GetUserInfo).Methods("GET")
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	var user *domain.User
	var userInfo um.UserInfoModel

	context := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		http.Error(w, e.ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	user, err = h.UserService.FindById(context, int(id))

	if err != nil {
		http.Error(w, e.ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, userError.ErrUserNotFound.Error(), http.StatusNoContent)
		return
	}

	errs := model.Copy(&userInfo, user)

	if errs != nil {
		http.Error(w, e.ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userInfo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
