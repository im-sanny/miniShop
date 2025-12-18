package repo

// Item represents a single item record
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

type ItemRepo interface {
	Get() ([]*Item, error)
	Create(i Item) (*Item, error)
	Update(i Item) (*Item, error)
	GetByID(itemID int) (*Item, error)
	Delete(itemID int) error
}

type itemRepo struct {
	itemList []*Item
}

// constructor or constructor function
func NewItemRepo() ItemRepo {
	repo := &itemRepo{}
	generateInitialItems(repo)
	return repo
}

// CreateItem adds a new item to ItemList and assigns it a new ID
func (r *itemRepo) Create(i Item) (*Item, error) {
	i.ID = len(r.itemList) + 1
	r.itemList = append(r.itemList, &i) // add item to the list
	return &i, nil                      // return the created item
}

// GetAllItem returns all items from ItemList
func (r *itemRepo) Get() ([]*Item, error) {
	return r.itemList, nil
}

// GetItemById finds an item by ID and returns a pointer to it
func (r *itemRepo) GetByID(itemID int) (*Item, error) {
	for _, item := range r.itemList { // loop through all items
		if item.ID == itemID { // check if ID matches
			return item, nil // return the matched item
		}
	}
	return nil, nil // return nil if item not found
}

// UpdateItem updates an existing item based on its ID
func (r *itemRepo) Update(item Item) (*Item, error) {
	for id, i := range r.itemList { // loop with index
		if i.ID == item.ID { // find matching ID
			r.itemList[id] = &item // replace old item with new one
		}
	}
	return &item, nil
}

// DeleteItemById removes an item from ItemList using its ID
func (r *itemRepo) Delete(itemID int) error {
	var filteredItem []*Item

	for _, i := range r.itemList { // loop through all items
		if i.ID != itemID { // keep items that do not match ID
			filteredItem = append(filteredItem, i)
		}
	}
	r.itemList = filteredItem // update ItemList after deletion
	return nil
}

// init seeds ItemList with some default data at startup
func generateInitialItems(r *itemRepo) {
	item1 := &Item{ID: 1, Name: "Blue Train", Brand: "John Coltrane", Price: 56.99}
	item2 := &Item{ID: 2, Name: "Jeru", Brand: "Gerry Mulligan", Price: 17.99}
	item3 := &Item{ID: 3, Name: "Sarah Vaughan and Clifford Brown", Brand: "Sarah Vaughan", Price: 39.99}

	r.itemList = append(r.itemList, item1, item2, item3) // add initial items
}
