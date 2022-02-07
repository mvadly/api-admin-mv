package user

import (
	"api-adminmv/config"
	"api-adminmv/entity"
	"api-adminmv/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var db = config.ConnectDB().Debug()

func ListUserBy(c *gin.Context) ([]entity.User, error) {
	users := []entity.User{}
	var page, _ = strconv.Atoi(c.Query("page"))
	var per_page, _ = strconv.Atoi(c.Query("per_page"))
	err := db.Select("id, username, fullname, status").
		Where("username LIKE ? OR fullname LIKE ? ", "%"+c.Query("q")+"%", "%"+c.Query("q")+"%").
		Order("created_at DESC").
		Offset((page - 1) * per_page).
		Limit(per_page).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func CountListUserBy(c *gin.Context) (int64, error) {
	users := []entity.User{}
	var count int64
	err := db.Select("id, username, fullname, status").
		Where("username LIKE ? OR fullname LIKE ? ", "%"+c.Query("q")+"%", "%"+c.Query("q")+"%").
		Order("created_at DESC").
		Find(&users).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func ListUser() ([]entity.User, error) {
	users := []entity.User{}
	err := db.Select("id, username, fullname, status").Order("created_at DESC").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func AddUser(req CreateReq) error {
	var users = &entity.User{
		Username:  req.Username,
		Password:  util.DefaultPassword(req.Username),
		Fullname:  req.Fullname,
		Status:    false,
		CreatedBy: "admin",
	}

	err := db.Create(users).Error
	return err
}

func UserUpdate(id string, req CreateReq) error {
	var users = &entity.User{
		Username:  req.Username,
		Fullname:  req.Fullname,
		Status:    false,
		UpdatedAt: time.Now(),
		UpdatedBy: "admin",
	}

	err := db.Where("id = ? ", id).Updates(users).Error
	return err
}

func UserDelete(id string) error {
	var users = &entity.User{}

	err := db.Where("id = ? ", id).Delete(users).Error
	return err
}
