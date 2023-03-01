package usecase

import (
	"context"

	"github.com/mahiro72/go-api-template/pkg/domain/model"
)

type User interface {
	Get(context.Context, *UserGetInput) (*UserGetOutput, error)
	Create(context.Context, *UserCreateInput) (*UserCreateOutput, error)
}

type UserGetInput struct {
	ID int
}
type UserGetOutput struct {
	User *model.User
}

type UserCreateInput struct {
	Name string
}
type UserCreateOutput struct {
	User *model.User
}
