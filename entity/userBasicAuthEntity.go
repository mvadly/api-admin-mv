package entity

type UserBasicAuth struct {
	ID       int    `gorm:"column:id;type:int(11);autoIncrement;primaryKey;" json:"id"`
	Username string `gorm:"column:username;type:varchar(20);not null:" json:"username"`
	Password string `gorm:"column:password;type:varchar(100);not null;" json:"password"`
	IsActive int    `gorm:"column:is_active;type:tinyint;default:1"`
}

func (UserBasicAuth) TableName() string {
	return "user_basic_auth"
}
