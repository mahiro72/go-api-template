package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/mahiro72/go-api-template/pkg/domain/model"
	"github.com/mahiro72/go-api-template/pkg/infra/dto"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (d *User) Get(ctx context.Context, id int) (*model.User, error) {
	return d.get(ctx, id)
}

func (d *User) Create(ctx context.Context, user *model.User) (*model.User, error) {
	dto := &dto.User{
		Name: user.Name,
	}

	query := "INSERT INTO users (name) VALUES (:name)"
	row, err := d.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		return nil, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}

	return d.get(ctx, (int)(id))
}

func (d *User) get(ctx context.Context, id int) (*model.User, error) {
	var dto dto.User

	query := "SELECT * FROM users WHERE id = ? LIMIT 1"

	err := d.db.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}

	return dto.ToModel(), nil
}
