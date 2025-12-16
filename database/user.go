package database

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (u User) StoreUser() User {
	if u.Id != 0 {
		return u
	}

	u.Id = len(users) + 1
	users = append(users, u)
	return u
}

func FindUser(email, pass string) *User {
	for index := range users {
		if users[index].Email == email && users[index].Password == pass {
			return &users[index]
		}
	}
	return nil
}
