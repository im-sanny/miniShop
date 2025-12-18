package repo

import "errors"

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var ErrUserNotFound = errors.New("user not found")

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

// func (u userRepo) Create(user User) (*User, error) {
// 	if user.Id != 0 {
// 		return &user, nil
// 	}

// 	user.Id = len(u.users) + 1
// 	u.users = append(u.users, user)
// 	return &user, nil
// }
func (u *userRepo) Create(user User) (*User, error) {
	user.Id = len(u.users) + 1
	u.users = append(u.users, user)
	return &user, nil
}

func (u *userRepo) Find(email, pass string) (*User, error) {
	for index := range u.users {
		if u.users[index].Email == email && u.users[index].Password == pass {
			return &u.users[index], nil
		}
	}
	return nil, ErrUserNotFound
}
