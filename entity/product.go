package entity

type Product struct {
	ID          int64   `gorm:"column:id;type:int(11) AUTO_INCREMENT;primaryKey" json:"id"`
	ProductName string  `gorm:"column:product_name;type:varchar(100);" json:"product_name"`
	Qty         int16   `gorm:"column:qty;type:int(11);" json:"qty"`
	QtyUnit     string  `gorm:"column:qty_unit;type:varchar(20);default:pcs;" json:"qty_unit"`
	Price       float32 `gorm:"column:price;type:double;" json:"price"`
	QtyPrice    string  `gorm:"column:qty_price;type:varchar(20);default:per pcs;" json:"qty_price"`
	Status      int     `gorm:"column:status;type:tinyint;default:0" json:"status"`
	ComodityId  string  `gorm:"column:comodity_id;type:varchar(30);" json:"comodity_id"`
	ImgProduct  string  `gorm:"column:img_product;type:text;" json:"img_product"`
	DateProduct string  `gorm:"column:date_product;type:date;" json:"date_product"`
	ExpProduct  string  `gorm:"column:exp_product;type:date;" json:"exp_product"`
	Global
}
