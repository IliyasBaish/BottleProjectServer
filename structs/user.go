package structs

type User struct {
	Id       int     `json:"id" db:"id"`
	Username string  `json:"username" binding:"required" db:"username"`
	Password string  `json:"password" binding:"required" db:"password_hash"`
	Wallet   float32 `json:"wallet" db:"wallet"`
	Role     string  `json:"role" db:"role"`
}
