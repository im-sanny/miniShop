package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Item represents a single item record
type Item struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Brand string  `json:"brand" db:"brand"`
	Price float64 `json:"price" db:"price"`
}

type ItemRepo interface {
	Create(i Item) (*Item, error)
	Get() ([]*Item, error)
	GetByID(itemID int) (*Item, error)
	Update(i Item) (*Item, error)
	Delete(itemID int) error
}

type itemRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
func NewItemRepo(db sqlx.DB) ItemRepo {
	return &itemRepo{
		db: &db,
	}
}

// CreateItem adds a new item to ItemList and assigns it a new ID
func (r *itemRepo) Create(i Item) (*Item, error) {
	query := `
		INSERT INTO items(
			name,
			brand,
			price
		) VALUES (
			$1,
			$2,
			$3
		)
		 RETURNING id
	`
	row := r.db.QueryRow(query, i.Name, i.Brand, i.Price)
	err := row.Scan(&i.ID)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

// GetAllItem returns all items from ItemList
func (r *itemRepo) Get() ([]*Item, error) {
	var itemList []*Item

	query := `
		SELECT
		id,
		name,
		brand,
		price
		from items
	`
	err := r.db.Select(&itemList, query)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

// GetItemById finds an item by ID and returns a pointer to it
func (r *itemRepo) GetByID(id int) (*Item, error) {
	var itm Item

	query := `
		SELECT
			id,
			name,
			brand,
			price
			from items
			where id =$1
	`
	err := r.db.Get(&itm, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &itm, nil
}

// UpdateItem updates an existing item based on its ID
func (r *itemRepo) Update(i Item) (*Item, error) {
	query := `
		UPDATE items
		SET name=$1, brand=$2, price=$3
		WHERE id=$4
	`
	row := r.db.QueryRow(query, i.Name, i.Brand, i.Price, i.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &i, nil
}

// DeleteItemById removes an item from ItemList using its ID
func (r *itemRepo) Delete(id int) error {
	query := `
		DELETE FROM items WHERE id =$1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
