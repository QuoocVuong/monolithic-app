package ginproduct

import (
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"

	"monolithic-app/common"
	"monolithic-app/modules/product/biz"
	//"monolithic-app/modules/product/model"
	"monolithic-app/modules/product/storage"
)

func ListItemGroup(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()

		store := storage.NewSqlStore(db)
		business := biz.NewListItemGroupBiz(store)

		// Gọi ListItemGroup (đã sửa) trong biz
		result, err := business.ListItemGroup(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Trả về response không có filter
		c.JSON(http.StatusOK, common.NewSuccessRespone(result, paging, nil))
	}
}
