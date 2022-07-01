package gateway

import (
	"context"
	"go_api_boilerplate/application/repository"
	"go_api_boilerplate/domain/entity"
	"go_api_boilerplate/infra/database"
)

func NewUserRepository(splHandler database.DBHandler) repository.UserRepository {
	return &userRepositoryImpl{
		sqlHandler: splHandler,
	}
}

type userRepositoryImpl struct {
	sqlHandler database.DBHandler
}

func (u userRepositoryImpl) FindAll(ctx context.Context) (*[]entity.User, error) {
	db, err := u.sqlHandler.Open(ctx)
	defer u.sqlHandler.Close(db)

	if err != nil {
		return nil, err
	}
	var users []entity.User
	db.Find(&users)
	return &users, err
}

func (u userRepositoryImpl) FindWithID(ctx context.Context, id int) (*entity.User, error) {
	db, err := u.sqlHandler.Open(ctx)
	defer u.sqlHandler.Close(db)

	if err != nil {
		return nil, err
	}
	var user entity.User
	db.Where("id = ?", id).First(&user)
	return &user, err
}

func (u userRepositoryImpl) Create(ctx context.Context, e entity.User) error {
	//TODO implement me
	panic("implement me")
}
