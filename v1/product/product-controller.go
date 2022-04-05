package product

import (
	"api-adminmv/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {

	data, count, err := GetListProduct(c)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err})
		return
	}

	var result = make(map[string]interface{})
	var page, _ = strconv.Atoi(c.Query("page"))
	var perpage, _ = strconv.Atoi(c.Query("perpage"))
	result["draw"] = c.Query("draw")
	result["recordsTotal"] = count
	result["recordsFiltered"] = len(data)
	result["page"] = page
	result["perpage"] = perpage
	result["data"] = data

	c.JSON(200, gin.H{"status": true, "message": "OK", "results": result})
}

func AddProduct(c *gin.Context) {
	var post Product
	// if err := c.ShouldBind(&post); err != nil {
	// 	c.JSON(400, gin.H{"err": err, "message": "error from bind struct"})
	// 	return
	// }
	c.ShouldBind(&post)

	err := post.Validate()
	if err != nil {
		c.JSON(400, gin.H{"status": false, "err": err})
		return
	}

	pathImg, str, err := util.UploadFile(c, "img_product", "images/")

	if err != nil {
		c.JSON(400, gin.H{"status": false, "err": err.Error(), "message": str})
		return
	}

	add, err := CreateProduct(post, pathImg)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "err": err})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "product saved",
		"result":  add,
	})

}
