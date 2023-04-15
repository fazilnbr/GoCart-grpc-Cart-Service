package domain

type Cart struct {
	Id      string `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	User_id string
}

type CartItem struct {
	Id         string `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Cart_id    string `json:"cartid" `
	Cart       *Cart  `json:"-" gorm:"foreignKey:Cart_id;references:Id"`
	Product_id string `json:"productid"`
	Quantity   int    `json:"quantity"`
}
