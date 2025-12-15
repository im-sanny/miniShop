package handlers

import (
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.GetAllItem(), http.StatusOK)
}
