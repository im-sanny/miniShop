package user

import (
	"encoding/json"
	"fmt"
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	createUser := newUser.StoreUser()
	util.SendData(w, createUser, http.StatusCreated)
}
