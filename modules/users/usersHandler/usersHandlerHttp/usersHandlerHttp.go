package usersHandlerHttp

import (
	"github.com/Rayato159/mytems-microservices/modules/users/usersUsecase"
)

type IUsersHandlerHttp interface{}

type usersHandlerHttp struct {
	usersUsecase usersUsecase.IUsersUsecase
}

func NewUsersHandlerHttp(usersUsecase usersUsecase.IUsersUsecase) IUsersHandlerHttp {
	return &usersHandlerHttp{
		usersUsecase: usersUsecase,
	}
}
