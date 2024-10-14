package ginproduct

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"monolithic-app/common"
	"monolithic-app/modules/product/biz"

	"monolithic-app/modules/product/storage"
	"net/http"
	"strconv"
)

func DeleteItemGroup(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSqlStore(db)
		business := biz.NewDeleteItemGroupBiz(store)
		if err := business.DeleteItemGroupById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(true))
	}
}
