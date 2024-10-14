package inventoryhandler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"monolithic-app/common"
	"monolithic-app/modules/inventory/biz"
	"monolithic-app/modules/inventory/model"
	"monolithic-app/modules/inventory/storage"
	"net/http"
	"strconv"
)

func CreateKhoHang(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.KhoHangCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewKhoHangBiz(store)

		if err := business.CreateNewKhoHang(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
func ListKhoHang(db *gorm.DB) func(*gin.Context) {
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
		business := biz.NewKhoHangBiz(store)

		result, err := business.ListKhoHang(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, common.NewSuccessRespone(result, paging, filter))
	}
}
func UpdateKhoHang(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // Lấy id từ URL
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		var data model.KhoHangUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewKhoHangBiz(store)

		// Gọi đến business logic để cập nhật dữ liệu
		if err := business.UpdateKhoHang(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
func DeleteKhoHang(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // Lấy id từ URL
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewKhoHangBiz(store)

		// Gọi đến business logic để cập nhật dữ liệu
		if err := business.DeleteKhoHang(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
