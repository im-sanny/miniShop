package util

type Claims struct {
	// standard claims
	Sub int64 `json:"sub"` // subject (use id)
	Exp int64 `json:"exp"` // expiration time (unix)
	Iat int64 `json:"iat"` // issued at
	Nbf int64 `json:"nbf"` // not valid before

	// custom claims
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}
