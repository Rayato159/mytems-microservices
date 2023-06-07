package itemsUsecase

import (
	"github.com/Rayato159/mytems-microservices/modules/items/itemsRepository"
)

type IItemsUsecase interface{}

type itemsUsecase struct {
	itemsRepository itemsRepository.IItemsRepository
}

func NewItemsUsecase(itemsRepository itemsRepository.IItemsRepository) IItemsUsecase {
	return &itemsUsecase{
		itemsRepository: itemsRepository,
	}
}
