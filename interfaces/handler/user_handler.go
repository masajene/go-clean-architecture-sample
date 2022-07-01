package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go_api_boilerplate/application/model"
	"go_api_boilerplate/application/usecase"
	"go_api_boilerplate/core"
	"go_api_boilerplate/domain/interactor"
	"go_api_boilerplate/infra/database"
	"go_api_boilerplate/interfaces/gateway"
	"net/http"
	"strconv"
)

var WireSet = wire.NewSet(
	NewUserHandler,
	interactor.NewUserUseCase,
	gateway.NewUserRepository,
	database.NewMysqlHandler,
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	handler := UserHandler{
		useCase: useCase,
	}
	return &handler
}

func (h *UserHandler) Users() gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := h.useCase.Users(c)
		if err != nil {
			core.BadRequest(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

func (h *UserHandler) User() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			core.BadRequest(c, err.Error())
			return
		}
		r, err := h.useCase.UserWithID(c, p)
		if err != nil {
			core.BadRequest(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, r)
	}
}
func (h *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h.useCase.CreateUser(c, model.UserModel{})
		if err != nil {
			core.BadRequest(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
