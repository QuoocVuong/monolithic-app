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

func UpdateItemGroup(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ItemGroupUpdate

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
		business := biz.NewUpdateItemGroupBiz(store)
		if err := business.UpdateItemGroupById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(true))
	}
}
