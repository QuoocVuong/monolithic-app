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

func CreateTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TonKho

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSqlStore(db)
		business := biz.NewTonKhoBiz(store)

		if err := business.CreateNewTonKho(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}

func ListTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()

		var filter model.Filterr
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewTonKhoBiz(store)

		result, err := business.ListTonKho(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, common.NewSuccessRespone(result, paging, filter))
	}
}

func UpdateTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // Lấy id từ URL
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		var data model.TonKho
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewTonKhoBiz(store)

		if err := business.UpdateTonKho(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}

func DeleteTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // Lấy id từ URL
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewTonKhoBiz(store)

		if err := business.DeleteTonKho(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
