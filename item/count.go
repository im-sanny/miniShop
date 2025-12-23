package item

func (svc *service) Count() (int64, error) {
	return svc.itemRepo.Count()
}
