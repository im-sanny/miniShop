package user

import (
	"encoding/json"
	"miniShop/config"
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
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

	cnf := config.GetConfig()
	accessToken, err := util.CreateSignedJwt(cnf.JWTSecretKey, util.Payload{
		Sub:       logUser.Id,
		FirstName: logUser.FirstName,
		LastName:  logUser.LastName,
		Email:     logUser.Email,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusOK)
}
