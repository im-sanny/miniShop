package database

var ItemList []Item

// Item represents a single item record
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

// CreateItem adds a new item to ItemList and assigns it a new ID
func CreateItem(i Item) Item {
	i.ID = len(ItemList) + 1
	ItemList = append(ItemList, i) // add item to the list
	return i                       // return the created item
}

// GetAllItem returns all items from ItemList
func GetAllItem() []Item {
	return ItemList
}

// GetItemById finds an item by ID and returns a pointer to it
func GetItemById(itemID int) *Item {
	for _, item := range ItemList { // loop through all items
		if item.ID == itemID { // check if ID matches
			return &item // return the matched item
		}
	}
	return nil // return nil if item not found
}

// UpdateItem updates an existing item based on its ID
func UpdateItem(item Item) {
	for id, i := range ItemList { // loop with index
		if i.ID == item.ID { // find matching ID
			ItemList[id] = item // replace old item with new one
		}
	}
}

// DeleteItemById removes an item from ItemList using its ID
func DeleteItemById(itemID int) {
	var filteredItem []Item

	for _, i := range ItemList { // loop through all items
		if i.ID != itemID { // keep items that do not match ID
			filteredItem = append(filteredItem, i)
		}
	}
	ItemList = filteredItem // update ItemList after deletion
}

// init seeds ItemList with some default data at startup
func init() {
	item1 := Item{ID: 1, Name: "Blue Train", Brand: "John Coltrane", Price: 56.99}
	item2 := Item{ID: 2, Name: "Jeru", Brand: "Gerry Mulligan", Price: 17.99}
	item3 := Item{ID: 3, Name: "Sarah Vaughan and Clifford Brown", Brand: "Sarah Vaughan", Price: 39.99}

	ItemList = append(ItemList, item1, item2, item3) // add initial items
}
