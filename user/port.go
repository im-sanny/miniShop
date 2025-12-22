package user

import (
	"miniShop/domain"
	userHandler "miniShop/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embedding
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
