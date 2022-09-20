package structs

type Deal struct {
	Id          int      `json:"id"`
	UserId      int      `json:"user_id"`
	Coins       float32  `json:"coins_value"`
	Date        string   `json:"-"`
	Bottles     []string `json:"bottles"`
	BottleCount int      `json:"bottle_count"`
}
