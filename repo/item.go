package repo

import (
	"database/sql"
	"miniShop/domain"
	"miniShop/item"

	"github.com/jmoiron/sqlx"
)

type ItemRepo interface {
	item.ItemRepo
}

type itemRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
func NewItemRepo(db *sqlx.DB) ItemRepo {
	return &itemRepo{
		db: db,
	}
}

// CreateItem adds a new item to ItemList and assigns it a new ID
func (r *itemRepo) Create(i domain.Item) (*domain.Item, error) {
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

// Get returns paginated items
func (r *itemRepo) Get(page, limit int64) ([]*domain.Item, error) {
	offset := (page - 1) * limit
	items := make([]*domain.Item, 0)

	query := `
		SELECT
		id,
		name,
		brand,
		price
		from items
		LIMIT $1 OFFSET $2
	`
	err := r.db.Select(&items, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *itemRepo) Count() (int64, error) {
	query := `SELECT COUNT(*)
	FROM items
	`
	var count int64
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetItemById finds an item by ID and returns a pointer to it
func (r *itemRepo) GetByID(id int) (*domain.Item, error) {
	var itm domain.Item

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
func (r *itemRepo) Update(i domain.Item) (*domain.Item, error) {
	query := `
		UPDATE items
		SET name=$1, brand=$2, price=$3
		WHERE id=$4
	`
	_, err := r.db.Exec(query, i.Name, i.Brand, i.Price, i.ID)
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
