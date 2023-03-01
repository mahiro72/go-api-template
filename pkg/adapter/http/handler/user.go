package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/mahiro72/go-api-template/pkg/adapter/http/response"
	"github.com/mahiro72/go-api-template/pkg/domain/model"
	"github.com/mahiro72/go-api-template/pkg/infra/registry"
	"github.com/mahiro72/go-api-template/pkg/interactor"
	"github.com/mahiro72/go-api-template/pkg/usecase"
)

type User struct {
	usecase usecase.User
}

func NewUser(r *registry.Repository) *User {
	usecase := interactor.NewUser(
		r.NewUser(),
	)
	return &User{usecase: usecase}
}

func (h *User) Get(w http.ResponseWriter, r *http.Request) {
	idParamString := chi.URLParam(r, "id")
	if idParamString == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	idParam, err := strconv.Atoi(idParamString)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.UserGetInput{
		ID: idParam,
	}

	out, err := h.usecase.Get(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newUserGetResponse(out.User)
	response.New(w, view)
}

func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.BadRequestErr(w, fmt.Errorf("error: content length is 0"))
		return
	}

	var j userCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.UserCreateInput{
		Name: j.Name,
	}

	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newUserCreateResponse(out.User)
	response.New(w, view)
}

type userCreateRequest struct {
	Name string `json:"name"`
}

type userCreateResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserCreateResponse(user *model.User) *userCreateResponse {
	return &userCreateResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}

type userGetResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserGetResponse(user *model.User) *userGetResponse {
	return &userGetResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}
