package item

import "miniShop/domain"

type service struct {
	itemRepo ItemRepo
}

func NewService(itemRepo ItemRepo) Service {
	return &service{
		itemRepo: itemRepo,
	}
}

func (svc *service) Create(item domain.Item) (*domain.Item, error) {
	return svc.itemRepo.Create(item)
}

func (svc *service) Get(page, limit int64) ([]*domain.Item, error) {
	return svc.itemRepo.Get(page, limit)
}

func (svc *service) GetByID(id int) (*domain.Item, error) {
	return svc.itemRepo.GetByID(id)
}

func (svc *service) Update(item domain.Item) (*domain.Item, error) {
	return svc.itemRepo.Update(item)
}

func (svc *service) Delete(id int) error {
	return svc.itemRepo.Delete(id)
}
