package producthandler

import (
	"errors"
	"github.com/gin-gonic/gin" // Gin framework cho web server
	"gorm.io/gorm"             // GORM ORM cho database
	"net/http"
	"strconv"

	"monolithic-app/common"                  // Các hàm và struct chung
	"monolithic-app/modules/product/biz"     // Logic nghiệp vụ cho product
	"monolithic-app/modules/product/model"   // Model cho product
	"monolithic-app/modules/product/storage" // Tương tác với database cho product
)

// CreateProduct là handler cho route POST /v1/products
// Tạo một sản phẩm mới.
func CreateProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ProductCreation // Struct chứa dữ liệu sản phẩm mới

		// Lấy dữ liệu JSON từ request body
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu dữ liệu không hợp lệ
			return
		}

		store := storage.NewSqlStore(db)           // Tạo đối tượng storage để tương tác với database
		business := biz.NewCreateProductBiz(store) // Tạo đối tượng biz để xử lý logic nghiệp vụ

		// Gọi hàm CreateNewProduct trong biz để tạo sản phẩm mới
		if err := business.CreateNewProduct(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu tạo sản phẩm thất bại
			return
		}

		// Trả về response thành công với ID của sản phẩm mới được tạo
		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(data.Id))
	}
}

// ListProduct là handler cho route GET /v1/products
// Lấy danh sách sản phẩm.
func ListProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging // Struct chứa thông tin phân trang

		// Lấy thông tin phân trang từ query parameters
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFullErrorResponse(common.ErrInvalidRequest(err)))
			return
		}
		paging.Process() // Xử lý thông tin phân trang

		var filter model.Filterr // Struct chứa thông tin lọc

		// Lấy thông tin lọc từ query parameters
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFullErrorResponse(common.ErrInvalidRequest(err)))
			return
		}

		store := storage.NewSqlStore(db)         // Tạo đối tượng storage
		business := biz.NewListProductBiz(store) // Tạo đối tượng biz

		// Gọi hàm ListProductById trong biz để lấy danh sách sản phẩm
		result, err := business.ListProductById(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewFullErrorResponse(err))
			return
		}

		// Trả về response thành công với danh sách sản phẩm, thông tin phân trang, và thông tin lọc
		c.JSON(http.StatusCreated, common.NewSuccessRespone(result, paging, filter))
	}
}

// GetProduct là handler cho route GET /v1/products/:id
// Lấy chi tiết một sản phẩm theo ID.
func GetProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Lấy ID từ URL parameters
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFullErrorResponse(common.ErrInvalidRequest(err)))
			return
		}

		store := storage.NewSqlStore(db)        // Tạo đối tượng storage
		business := biz.NewGetProductBiz(store) // Tạo đối tượng biz

		// Gọi hàm GetProductById trong biz để lấy chi tiết sản phẩm
		data, err := business.GetProductById(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) { // Kiểm tra nếu sản phẩm không tồn tại
				c.JSON(http.StatusNotFound, common.NewFullErrorResponse(common.ErrCannotGetEntity("sản phẩm", err))) // Trả về 404 Not Found
				return
			}
			c.JSON(http.StatusInternalServerError, common.NewFullErrorResponse(err)) // Xử lý các lỗi khác
			return
		}
		// Trả về TOÀN BỘ OBJECT SẢN PHẨM, không chỉ ID
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(data)) // Sử dụng 200 OK cho GET request thành công
	}
}

// UpdateProduct là handler cho route PATCH /v1/products/:id
// Cập nhật thông tin một sản phẩm theo ID.
func UpdateProduct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.ProductUpdate // Struct chứa thông tin cập nhật sản phẩm

		// Lấy ID từ URL parameters
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Lấy dữ liệu JSON từ request body
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSqlStore(db)           // Tạo đối tượng storage
		business := biz.NewUpdateProductBiz(store) // Tạo đối tượng biz

		// Gọi hàm UpdateProductById trong biz để cập nhật sản phẩm
		if err := business.UpdateProductById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Trả về response thành công
		c.JSON(http.StatusCreated, common.SimpleSuccessRespone(true))
	}
}

// DeleteProduct là handler cho route DELETE /v1/products/:id
// Xóa một sản phẩm theo ID.
//
//	func DeleteProduct(db *gorm.DB) func(*gin.Context) {
//		return func(c *gin.Context) {
//			// Lấy ID từ URL parameters
//			id, err := strconv.Atoi(c.Param("id"))
//			if err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//				return
//			}
//
//			store := storage.NewSqlStore(db)           // Tạo đối tượng storage
//			business := biz.NewDeleteProductBiz(store) // Tạo đối tượng biz
//
//			// Gọi hàm DeleteProductById trong biz để xóa sản phẩm
//			if err := business.DeleteProductById(c.Request.Context(), id); err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//				return
//			}
//
//			// Trả về response thành công
//			c.JSON(http.StatusCreated, common.SimpleSuccessRespone(true))
//		}
//	}
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		statusDeleted := model.ProductStatusDeleted // Hoặc giá trị tương ứng

		// Tìm sản phẩm trước để kiểm tra xem nó có tồn tại không
		var product model.SanPham
		if err := db.First(&product, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Kiểm tra xem sản phẩm đã bị xóa chưa
		if product.Status != nil && *product.Status == statusDeleted {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product already deleted"})
			return
		}

		if err := db.Model(&model.SanPham{}).Where("id = ?", id).Update("status", &statusDeleted).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
