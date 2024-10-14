package inventoryhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" // Gin framework cho web server
	"gorm.io/gorm"             // GORM ORM cho database

	"monolithic-app/common"                    // Các hàm và struct chung
	"monolithic-app/modules/inventory/biz"     // Logic nghiệp vụ cho inventory
	"monolithic-app/modules/inventory/model"   // Model cho inventory
	"monolithic-app/modules/inventory/storage" // Tương tác với database cho inventory
)

// CreateDuKienTonKho là handler cho route POST /v1/du-kien-ton-khos
// Tạo một dự kiến tồn kho mới.
func CreateDuKienTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.DuKienTonKho // Struct chứa dữ liệu dự kiến tồn kho mới

		// Lấy dữ liệu từ request body
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu dữ liệu không hợp lệ
			return
		}

		store := storage.NewSqlStore(db)          // Tạo đối tượng storage
		business := biz.NewDuKienTonKhoBiz(store) // Tạo đối tượng biz

		// Gọi hàm CreateNewDuKienTonKho trong biz để tạo dự kiến tồn kho mới
		if err := business.CreateNewDuKienTonKho(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu tạo thất bại
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true)) // Trả về response thành công
	}
}

// ListDuKienTonKho là handler cho route GET /v1/du-kien-ton-khos
// Lấy danh sách dự kiến tồn kho.
func ListDuKienTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging // Struct chứa thông tin phân trang

		// Lấy thông tin phân trang từ query parameters
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process() // Xử lý thông tin phân trang

		var filter model.Filterr // Struct chứa thông tin lọc

		// Lấy thông tin lọc từ query parameters
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSqlStore(db)          // Tạo đối tượng storage
		business := biz.NewDuKienTonKhoBiz(store) // Tạo đối tượng biz

		// Gọi hàm ListDuKienTonKho trong biz để lấy danh sách dự kiến tồn kho
		result, err := business.ListDuKienTonKho(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu lấy danh sách thất bại
			return
		}

		// Trả về response thành công với danh sách dự kiến tồn kho, thông tin phân trang, và thông tin lọc
		c.JSON(http.StatusCreated, common.NewSuccessRespone(result, paging, filter))
	}
}

// UpdateDuKienTonKho là handler cho route PATCH /v1/du-kien-ton-khos/:id
// Cập nhật thông tin một dự kiến tồn kho theo ID.
func UpdateDuKienTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Lấy ID từ URL parameters
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		var data model.DuKienTonKho // Struct chứa dữ liệu dự kiến tồn kho cần cập nhật

		// Lấy dữ liệu từ request body
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSqlStore(db)          // Tạo đối tượng storage
		business := biz.NewDuKienTonKhoBiz(store) // Tạo đối tượng biz

		// Gọi hàm UpdateDuKienTonKho trong biz để cập nhật dự kiến tồn kho
		if err := business.UpdateDuKienTonKho(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu cập nhật thất bại
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true)) // Trả về response thành công
	}
}

// DeleteDuKienTonKho là handler cho route DELETE /v1/du-kien-ton-khos/:id
// Xóa một dự kiến tồn kho theo ID.
func DeleteDuKienTonKho(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Lấy ID từ URL parameters
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		store := storage.NewSqlStore(db)          // Tạo đối tượng storage
		business := biz.NewDuKienTonKhoBiz(store) // Tạo đối tượng biz

		// Gọi hàm DeleteDuKienTonKho trong biz để xóa dự kiến tồn kho
		if err := business.DeleteDuKienTonKho(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu xóa thất bại
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true)) // Trả về response thành công
	}
}
