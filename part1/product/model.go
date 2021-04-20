package product

type Product struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" form:"name"`
	Cost       int    `json:"cost" form:"cost"`
	Price      int    `json:"price" form:"price"`
	Note       string `json:"note" form:"note"`
	Is_sell    bool   `json:"is_sell" form:"is_sell"`
	Start_time string `json:"start_time" form:"start_time"`
	End_time   string `json:"end_time" form:"end_time"`
}

func (Product) TableName() string {
	return "product"
}
