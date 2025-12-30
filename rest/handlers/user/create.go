package user

import (
	"encoding/json"
	"miniShop/domain"
	"miniShop/util"
	"net/http"
	"strings"
)

type reqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req reqCreateUser
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		util.SendError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	usr, err := h.svc.Create(domain.User{
		FirstName:   strings.TrimSpace(req.FirstName),
		LastName:    strings.TrimSpace(req.LastName),
		Email:       strings.TrimSpace(req.Email),
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	util.SendData(w, http.StatusCreated, usr)
}
