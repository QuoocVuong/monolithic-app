package itemgrouphandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"monolithic-app/common"
	"monolithic-app/modules/itemgroup/biz"
	"monolithic-app/modules/itemgroup/model"
	"monolithic-app/modules/itemgroup/storage"
)

// CreateItemGroup ...
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

// ListItemGroup ...
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

// GetItemGroup ...
func GetItemGroup(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSqlStore(db)
		business := biz.NewGetItemGroupBiz(store)

		data, err := business.GetItemGroupById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRespone(data)) // Trả về data thay vì data.Id
	}
}

// UpdateItemGroup ...
func UpdateItemGroup(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ItemGroupUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewUpdateItemGroupBiz(store)
		if err := business.UpdateItemGroupById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(true))
	}
}

// DeleteItemGroup ...
func DeleteItemGroup(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSqlStore(db)
		business := biz.NewDeleteItemGroupBiz(store)

		if err := business.DeleteItemGroupById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(true))
	}
}
