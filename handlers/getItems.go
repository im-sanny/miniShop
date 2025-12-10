package handlers

import (
	"miniShop/database"
	"miniShop/util"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ItemList, 200)
}
