package ginproduct

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"monolithic-app/common"
	"monolithic-app/modules/product/biz"
	"monolithic-app/modules/product/model"
	"monolithic-app/modules/product/storage"
	"net/http"
	"strconv"
)

func UpdateProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ProductUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewUpdateProductBiz(store)
		if err := business.UpdateProductById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(true))
	}
}
