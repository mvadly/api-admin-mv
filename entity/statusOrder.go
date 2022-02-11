package entity

type StatusOrder struct {
	ID         int    `gorm:"column:id;type:int(5) AUTO_INCREMENT;primaryKey" json:"id"`
	StatusName string `gorm:"column:status_name;type:varchar(20);" json:"status_name"`
	Global
}
