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

func CreateItemGroup(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ItemGroupCreation
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewCreateItemGroupBiz(store)
		if err := business.CreateNewItemGroup(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(data.Id))
	}
}
