package product

import (
	"api-adminmv/config"
	"api-adminmv/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db = config.ConnectDB().Debug()

func selectProduct(c *gin.Context) *gorm.DB {
	var search = c.Query("search")

	query := db.Select("id, product_name, qty, price, status, comodity_id, created_at").
		Where("product_name LIKE ?", "%"+search+"%")
	return query

}

func GetListProduct(c *gin.Context) ([]entity.Product, int64, error) {
	var products []entity.Product
	var page, _ = strconv.Atoi(c.Query("page"))
	var perpage, _ = strconv.Atoi(c.Query("perpage"))
	var start = (page - 1) * perpage
	var count int64
	query := selectProduct(c).Offset(start).Limit(perpage).Order("created_at desc, updated_at")
	err := query.Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Find(&products).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return products, count, nil

}

func CreateProduct(p Product, pathImage interface{}) (entity.Product, error) {
	var product = entity.Product{
		ProductName: p.ProductName,
		Qty:         p.Qty,
		QtyUnit:     p.QtyUnit,
		Price:       p.Price,
		QtyPrice:    p.QtyPrice,
		ComodityId:  p.ComodityId,
		ImgProduct:  pathImage.(string),
		Global: entity.Global{
			CreatedBy: "admin",
		},
	}

	query := db.Create(&product)
	if query.Error != nil || query.RowsAffected == 0 {
		return entity.Product{}, query.Error
	}

	return product, nil
}
