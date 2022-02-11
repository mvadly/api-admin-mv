package entity

import "time"

type Global struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy string    `gorm:"column:created_by;type:varchar(50) not null" json:"created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	UpdatedBy string    `gorm:"column:updated_by;type:varchar(50)" json:"updated_by"`
}
