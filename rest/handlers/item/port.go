package item

import (
	"miniShop/domain"
)

type Service interface {
	Create(i domain.Item) (*domain.Item, error)
	Get(page, limit int64) ([]*domain.Item, error)
	GetByID(itemID int) (*domain.Item, error)
	Update(i domain.Item) (*domain.Item, error)
	Delete(itemID int) error
	Count() (int64, error)
}
