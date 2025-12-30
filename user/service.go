package user

import (
	"miniShop/domain"
	"strings"
)

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	return svc.userRepo.Create(user)
}

func (svc *service) Find(email string, pass string) (*domain.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	return svc.userRepo.Find(email, pass)
}
