package model

type Shop struct {
	Id            int64  `xorm:"pk autoincr" json:"id"`
	Name          string `xorm:"varchar(12)" json:"name"`
	PromotionInfo string `xorm:"varchar(30)" json:"promotion_info"`
	Address       string `xorm:"varchar(100)" json:"address"`
	Phone         string `xorm:"varchar(11)" json:"phone"`
	Status        int    `xorm:"tinyint" json:"status"`

	Longitude float64 `xorm:"" json:"longitude"`
	Latitude  float64 `xorm:"" json:"latitude"`

	IsNew     bool `xorm:"bool" json:"is_new"`
	IsPremium bool `xorm:"bool" json:"is_premium"`

	Rating      float32 `xorm:"float" json:"rating"`
	RatingCount int64   `xorm:"int" json:"rating_count"`

	RecentOrderNum int64 `xorm:"int" json:"recent_order_num"`
}
