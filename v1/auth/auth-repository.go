package auth

import (
	"api-adminmv/config"
	"api-adminmv/entity"
)

var db = config.ConnectDB().Debug()

func GetUserPasswordLogin(req RequestLogin) (entity.User, error) {
	users := entity.User{}
	get := db.Select("id, username, password, fullname, status").
		Where("username = ? ", req.Username).
		Find(&users)
	if get.Error != nil {
		return users, get.Error
	}
	return users, nil

}
