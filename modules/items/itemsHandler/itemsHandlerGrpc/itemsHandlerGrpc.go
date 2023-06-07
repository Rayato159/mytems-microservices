package itemsHandlerGrpc

import (
	"github.com/Rayato159/mytems-microservices/modules/items/itemsUsecase"
)

type IItemsHandlerGrpc interface{}

type itemsHandlerGrpc struct {
	itemsUsecase itemsUsecase.IItemsUsecase
}

func NewItemsHandlerGrpc(itemsUsecase itemsUsecase.IItemsUsecase) IItemsHandlerGrpc {
	return &itemsHandlerGrpc{
		itemsUsecase: itemsUsecase,
	}
}
