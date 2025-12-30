package user

import (
	"encoding/json"
	"miniShop/util"
	"net/http"
	"strings"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req loginReq
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	if req.Email == "" || req.Password == "" {
		util.SendError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	accessToken, err := util.CreateSignedJwt(h.cnf.JWTSecretKey, util.Payload{
		Sub: usr.ID,
		// FirstName: usr.FirstName,
		// LastName:  usr.LastName,
		Email: usr.Email,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusOK, accessToken)
}
