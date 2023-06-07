package itemsHandlerHttp

import (
	"github.com/Rayato159/mytems-microservices/modules/items/itemsUsecase"
)

type IItemsHandlerHttp interface{}

type itemsHandlerHttp struct {
	itemsUsecase itemsUsecase.IItemsUsecase
}

func NewItemsHandlerHttp(itemsUsecase itemsUsecase.IItemsUsecase) IItemsHandlerHttp {
	return &itemsHandlerHttp{
		itemsUsecase: itemsUsecase,
	}
}
