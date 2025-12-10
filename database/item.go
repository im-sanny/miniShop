package database

var ItemList []Item

// album represents data about a record album
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

// albums slice to seed record album data.
func init() {
	item1 := Item{ID: 1, Name: "Blue Train", Brand: "John Coltrane", Price: 56.99}
	item2 := Item{ID: 2, Name: "Jeru", Brand: "Gerry Mulligan", Price: 17.99}
	item3 := Item{ID: 3, Name: "Sarah Vaughan and Clifford Brown", Brand: "Sarah Vaughan", Price: 39.99}

	ItemList = append(ItemList, item1, item2, item3)
}
