package domain

// Item represents a single item record
type Item struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Brand string  `json:"brand" db:"brand"`
	Price float64 `json:"price" db:"price"`
	Image string  `json:"image" db:"image"`
}
