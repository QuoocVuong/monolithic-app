package ginproduct

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"monolithic-app/common"
	"monolithic-app/modules/product/biz"
	"monolithic-app/modules/product/model"
	"monolithic-app/modules/product/storage"
	"net/http"
)

func CreateProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ProductCreation
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.CreateNewProductBiz(store)
		if err := business.CreateNewProduct(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(data.Id))
	}
}
