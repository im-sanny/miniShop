package item

import (
	"miniShop/domain"
	itemHandler "miniShop/rest/handlers/item"
)

type Service interface {
	itemHandler.Service
}

type ItemRepo interface {
	Create(i domain.Item) (*domain.Item, error)
	Get(page, limit int64) ([]*domain.Item, error)
	GetByID(itemID int) (*domain.Item, error)
	Update(i domain.Item) (*domain.Item, error)
	Delete(itemID int) error
	Count() (int64, error)
}
