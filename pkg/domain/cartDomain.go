package domain

type Cart struct {
	Id      int64 `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	User_id int64 `json:"userid"`
}

type CartItem struct {
	Id         int64 `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	Cart_id    int64 `json:"-" `
	Cart       *Cart `json:"-" gorm:"foreignKey:Cart_id;references:Id"`
	Product_id int64 `json:"productid"`
	Quantity   int64 `json:"quantity"`
}
