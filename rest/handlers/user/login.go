package user

import (
	"encoding/json"
	"fmt"
	"miniShop/util"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	accessToken, err := util.CreateSignedJwt(h.cnf.JWTSecretKey, util.Payload{
		Sub:       usr.Id,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusOK, accessToken)
}
