package user

import (
	"encoding/json"
	"fmt"
	"miniShop/repo"
	"miniShop/util"
	"net/http"
)

type ReqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser ReqCreateUser

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Create(repo.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
	}

	util.SendData(w, http.StatusCreated, usr)
}
