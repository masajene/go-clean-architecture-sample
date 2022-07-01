package usecase

import (
	"context"
	"go_api_boilerplate/application/model"
)

type UserUseCase interface {
	Users(ctx context.Context) (*[]model.UserModel, error)
	UserWithID(ctx context.Context, id int) (*model.UserModel, error)
	CreateUser(ctx context.Context, m model.UserModel) error
}
