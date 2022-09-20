package structs

type UnclaimedDeal struct {
	Id          int      `json:"id"`
	UserId      int      `json:"user_id"`
	Coins       float32  `json:"coins_value"`
	Date        string   `json:"-"`
	Bottles     []string `json:"bottles"`
	BottleCount int      `json:"bottle_count"`
	Station     int      `json:"station"`
	Time        string   `json:"-"`
	Verified    string   `json:"-"`
}
