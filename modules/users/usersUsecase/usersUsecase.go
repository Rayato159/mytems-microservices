package usersUsecase

import (
	"github.com/Rayato159/mytems-microservices/modules/users/usersRepository"
)

type IUsersUsecase interface{}

type usersUsecase struct {
	usersRepository usersRepository.IUsersRepository
}

func NewUsersUsecase(usersRepository usersRepository.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		usersRepository: usersRepository,
	}
}
