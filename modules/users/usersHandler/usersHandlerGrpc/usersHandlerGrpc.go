package usersHandlerGrpc

import (
	"github.com/Rayato159/mytems-microservices/modules/users/usersUsecase"
)

type IUsersHandlerGrpc interface{}

type usersHandlerGrpc struct {
	usersUsecase usersUsecase.IUsersUsecase
}

func NewUsersHandlerGrpc(usersUsecase usersUsecase.IUsersUsecase) IUsersHandlerGrpc {
	return &usersHandlerGrpc{
		usersUsecase: usersUsecase,
	}
}
