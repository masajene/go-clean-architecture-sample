package repository

import (
	"context"
	"go_api_boilerplate/domain/entity"
)

type UserRepository interface {
	FindAll(ctx context.Context) (*[]entity.User, error)
	FindWithID(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, e entity.User) error
}
