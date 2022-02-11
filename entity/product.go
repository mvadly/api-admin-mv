package entity

type Product struct {
	ID          int64   `gorm:"column:id;type:int(11) AUTO_INCREMENT;primaryKey" json:"id"`
	ProductName string  `gorm:"column:product_name;type:varchar(100);" json:"product_name"`
	Qty         int16   `gorm:"column:qty;type:int(11);" json:"qty"`
	Price       float32 `gorm:"column:price;type:double;" json:"price"`
	Status      int     `gorm:"column:status;type:tinyint;" json:"status"`
	ComodityId  int     `gorm:"column:comodity_id;type:int(5);" json:"comodity_id"`
	Global
}
