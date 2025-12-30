package repo

import (
	"database/sql"
	"miniShop/domain"
	"miniShop/user"
	"miniShop/util"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// UserRepo defines the repository interface exposed by this package.
// It embeds the domain-level UserRepo contract.
type UserRepo interface {
	user.UserRepo
}

// userRepo is the concrete database-backed implementation
// of the UserRepo interface.
type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo creates a new User repository instance.
// It returns the interface type to enforce dependency inversion.
func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

// Create persists a new user in the database.
//
// Responsibilities:
// - Hash the user's password before storing
// - Insert user data into the database
// - Retrieve and assign the generated user ID
//
// This method does NOT:
// - Validate business rules
// - Handle HTTP concerns
func (u *userRepo) Create(user domain.User) (*domain.User, error) {
	// Hash plaintext password using bcrypt
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	// Replace plaintext password with hashed value
	user.Password = string(hashed)

	// SQL query to insert user and return generated ID
	query := `
	INSERT INTO users (
		first_name,
		last_name,
		email,
		password,
		is_shop_owner
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	// Execute insert and scan generated ID
	err = u.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.IsShopOwner,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Find retrieves a user by email and validates credentials.
//
// Responsibilities:
// - Fetch user record by email
// - Compare stored password hash with provided password
// - Return a generic authentication error on failure
//
// Security notes:
// - Password comparison is done in Go, not SQL
// - Same error is returned for invalid email or password
func (u *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User

	// Query user by email only
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1
	LIMIT 1
	`

	err := u.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			// Do not reveal whether email exists
			return nil, util.ErrInvalidCredentials
		}
		return nil, err
	}

	// Compare stored hash with provided password
	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(pass),
	) != nil {
		return nil, util.ErrInvalidCredentials
	}

	return &user, nil
}
