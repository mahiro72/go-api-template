package dto

import (
	"time"

	"github.com/mahiro72/go-api-template/pkg/domain/model"
)

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func NewUserDtoFromModel(m *model.User) *User {
	return &User{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}

func (d *User) ToModel() *model.User {
	return &model.User{
		ID:        d.ID,
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
	}
}
