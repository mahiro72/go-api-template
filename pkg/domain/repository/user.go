package repository

import (
	"context"

	"github.com/mahiro72/go-api-template/pkg/domain/model"
)

type User interface {
	Get(context.Context, int) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
}
