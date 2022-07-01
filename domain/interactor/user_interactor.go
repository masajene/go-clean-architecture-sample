package interactor

import (
	"context"
	"go_api_boilerplate/application/model"
	"go_api_boilerplate/application/repository"
	"go_api_boilerplate/application/usecase"
	"go_api_boilerplate/domain/entity"
	"go_api_boilerplate/domain/validation"
	"time"
)

type userUseCaseImpl struct {
	repository repository.UserRepository
}

func NewUserUseCase(
	repository repository.UserRepository) usecase.UserUseCase {
	return &userUseCaseImpl{
		repository: repository,
	}
}

func (u userUseCaseImpl) Users(ctx context.Context) (*[]model.UserModel, error) {
	e, err := u.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var es []model.UserModel
	for _, v := range *e {
		es = append(es, model.UserModel{
			ID:          v.ID,
			Name:        v.Name,
			MailAddress: v.MailAddress,
			Birthday:    v.Birthday,
		})
	}
	return &es, err
}

func (u userUseCaseImpl) UserWithID(ctx context.Context, id int) (*model.UserModel, error) {
	e, err := u.repository.FindWithID(ctx, id)
	if err != nil {
		return nil, err
	}
	m := model.UserModel{
		ID:          e.ID,
		Name:        e.Name,
		MailAddress: e.MailAddress,
		Birthday:    e.Birthday,
	}
	return &m, err
}

func (u userUseCaseImpl) CreateUser(ctx context.Context, m model.UserModel) error {
	vl, err := validation.NewValidation()
	err = vl.Struct(m)
	if err != nil {
		return err
	}
	e := entity.User{
		ID:          m.ID,
		Name:        m.Name,
		MailAddress: m.MailAddress,
		Birthday:    m.Birthday,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	return u.repository.Create(ctx, e)
}
