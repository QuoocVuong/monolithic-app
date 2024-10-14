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

func ListProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		paging.Process()
		var filter model.Filterr
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewListProductBiz(store)
		result, err := business.ListProductById(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, common.NewSuccessRespone(result, paging, filter))
	}
}
