package handlers

import (
	"miniShop/database"
	"miniShop/util"
	"net/http"
	"strconv"
)

func GetItemById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("itemId") // All values coming from the URL are strings. So we collect the itemId as a string first, then convert it to int so database queries can use it.

	id, err := strconv.Atoi(idStr) // converting idStr into int which was in string form then it'll store in id
	if err != nil {                // id error qual not nil then it has error
		http.Error(w, "give me valid item id", http.StatusBadRequest) // after getting error this line will return an error response
		return                                                        // from here req will go back or stop
	}

	for _, item := range database.ItemList { // this will look at the range of the database for getting id later
		if item.ID == id { // if item id == id then
			util.SendData(w, item, http.StatusOK) // this will encode the item as JSON and send it in the response
			return                                // return whatever it got
		}
	}

	util.SendData(w, "item not found", http.StatusNotFound) //if no item with that ID is found it will throw error
}
