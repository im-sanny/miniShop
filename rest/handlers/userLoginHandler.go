package handlers

import (
	"encoding/json"
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginReq

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	logUser := database.FindUser(loginReq.Email, loginReq.Password)
	if logUser == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	util.SendData(w, logUser, http.StatusOK)
}
