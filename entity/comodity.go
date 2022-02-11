package entity

type Comodity struct {
	ID           int    `gorm:"column:id;type:int(5) AUTO_INCREMENT;primaryKey" json:"id"`
	ComodityName string `gorm:"column:comodity_name;type:varchar(20);" json:"comodity_name"`
	Global
}
