package interactor

import (
	"context"
	"fmt"

	"github.com/mahiro72/go-api-template/pkg/domain/model"
	"github.com/mahiro72/go-api-template/pkg/domain/repository"
	"github.com/mahiro72/go-api-template/pkg/usecase"
)

type User struct {
	repositroy repository.User
}

func NewUser(repository repository.User) *User {
	return &User{repositroy: repository}
}

func (uc *User) Get(ctx context.Context, in *usecase.UserGetInput) (*usecase.UserGetOutput, error) {
	if in.ID == 0 {
		return nil, fmt.Errorf("id is invalid")
	}

	user, err := uc.repositroy.Get(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	return &usecase.UserGetOutput{
		User: user,
	}, nil
}

func (uc *User) Create(ctx context.Context, in *usecase.UserCreateInput) (*usecase.UserCreateOutput, error) {
	if in.Name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	if len(in.Name) > 100 {
		return nil, fmt.Errorf("name must be 100 characters or less in length")
	}

	u := &model.User{
		Name: in.Name,
	}

	user, err := uc.repositroy.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return &usecase.UserCreateOutput{
		User: user,
	}, nil
}
