package entity

import "time"

type User struct {
	ID        int       `gorm:"column:id;type:int(11) not null;autoIncrement;unique;primaryKey;" json:"id"`
	Username  string    `gorm:"column:username;type:varchar(50) not null;unique" json:"username"`
	Password  string    `gorm:"column:password;type:varchar(255) not null" json:"password"`
	Fullname  string    `gorm:"column:fullname;type:varchar(100) not null" json:"fullname"`
	Status    bool      `gorm:"column:status;type:varchar(50) not null" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy string    `gorm:"column:created_by;type:varchar(50) not null" json:"created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	UpdatedBy string    `gorm:"column:updated_by;type:varchar(50)" json:"updated_by"`
}
