package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"monolithic-app/modules/inventory/transport/inventoryhandler"
	"monolithic-app/modules/itemgroup/transport/itemgrouphandler"
	"monolithic-app/modules/product/model"
	"monolithic-app/modules/product/transport/producthandler"
)

func main() {
	// Database setup
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" // Chuỗi kết nối database MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})                                  // Mở kết nối đến database
	if err != nil {
		log.Fatal(err) // Nếu lỗi, in lỗi và dừng chương trình
	}

	// Tự động migrate các model vào database để tạo bảng
	db.AutoMigrate(
		&model.NhomHang{},     // Tạo bảng nhóm hàng
		&model.KhoHang{},      // Tạo bảng kho hàng
		&model.SanPham{},      // Tạo bảng sản phẩm
		&model.TonKho{},       // Tạo bảng tồn kho
		&model.DuKienTonKho{}, // Tạo bảng dự kiến tồn kho
	)

	// Khởi tạo Gin router
	r := gin.Default()

	// Tạo group /v1 cho API
	v1 := r.Group("/v1")
	{
		// ======================================= PRODUCT =======================================
		// Routes cho product
		products := v1.Group("/products") // Tạo group /v1/products
		{
			products.POST("", producthandler.CreateProduct(db))       // POST /v1/products: Tạo sản phẩm mới
			products.GET("", producthandler.ListProduct(db))          // GET /v1/products: Lấy danh sách sản phẩm
			products.PATCH("/:id", producthandler.UpdateProduct(db))  // PATCH /v1/products/:id: Cập nhật sản phẩm
			products.DELETE("/:id", producthandler.DeleteProduct(db)) // DELETE /v1/products/:id: Xóa sản phẩm
			products.GET("/:id", producthandler.GetProduct(db))       // GET /v1/products/:id: Lấy chi tiết sản phẩm
		}

		// ======================================= ITEM GROUP =======================================
		// Routes cho itemgroup
		itemgroups := v1.Group("/itemgroups") // Tạo group /v1/itemgroups
		{
			itemgroups.POST("", itemgrouphandler.CreateItemGroup(db))       // POST /v1/itemgroups: Tạo nhóm hàng mới
			itemgroups.GET("", itemgrouphandler.ListItemGroup(db))          // GET /v1/itemgroups: Lấy danh sách nhóm hàng
			itemgroups.PATCH("/:id", itemgrouphandler.UpdateItemGroup(db))  // PATCH /v1/itemgroups/:id: Cập nhật nhóm hàng
			itemgroups.DELETE("/:id", itemgrouphandler.DeleteItemGroup(db)) // DELETE /v1/itemgroups/:id: Xóa nhóm hàng
			itemgroups.GET("/:id", itemgrouphandler.GetItemGroup(db))       // GET /v1/itemgroups/:id: Lấy chi tiết nhóm hàng
		}

		// ======================================= INVENTORY =======================================
		// Routes cho kho_hang
		khohangs := v1.Group("/khohangs") // Tạo group /v1/kho-hangs
		{
			khohangs.POST("", inventoryhandler.CreateKhoHang(db))       // POST /v1/kho-hangs: Tạo kho hàng mới
			khohangs.GET("", inventoryhandler.ListKhoHang(db))          // GET /v1/kho-hangs: Lấy danh sách kho hàng
			khohangs.PATCH("/:id", inventoryhandler.UpdateKhoHang(db))  // PATCH /v1/kho-hangs/:id: Cập nhật kho hàng
			khohangs.DELETE("/:id", inventoryhandler.DeleteKhoHang(db)) // DELETE /v1/kho-hangs/:id: Xóa kho hàng
			//khohangs.GET("/:id", inventoryhandler.GetKhoHang(db))      // GET /v1/kho-hangs/:id: Lấy chi tiết kho hàng
		}

		// Routes cho ton_kho
		tonkhos := v1.Group("/tonkhos") // Tạo group /v1/ton-khos
		{
			tonkhos.POST("", inventoryhandler.CreateTonKho(db))       // POST /v1/ton-khos: Tạo tồn kho mới
			tonkhos.GET("", inventoryhandler.ListTonKho(db))          // GET /v1/ton-khos: Lấy danh sách tồn kho
			tonkhos.PATCH("/:id", inventoryhandler.UpdateTonKho(db))  // PATCH /v1/ton-khos/:id: Cập nhật tồn kho
			tonkhos.DELETE("/:id", inventoryhandler.DeleteTonKho(db)) // DELETE /v1/ton-khos/:id: Xóa tồn kho
			//tonkhos.GET("/:id", inventoryhandler.GetTonKho(db))      // GET /v1/ton-khos/:id: Lấy chi tiết tồn kho
		}

		// Routes cho du_kien_ton_kho
		dukienTonkhos := v1.Group("/dukientonkhos") // Tạo group /v1/du-kien-ton-khos
		{
			dukienTonkhos.POST("", inventoryhandler.CreateDuKienTonKho(db))       // POST /v1/du-kien-ton-khos: Tạo dự kiến tồn kho mới
			dukienTonkhos.GET("", inventoryhandler.ListDuKienTonKho(db))          // GET /v1/du-kien-ton-khos: Lấy danh sách dự kiến tồn kho
			dukienTonkhos.PATCH("/:id", inventoryhandler.UpdateDuKienTonKho(db))  // PATCH /v1/du-kien-ton-khos/:id: Cập nhật dự kiến tồn kho
			dukienTonkhos.DELETE("/:id", inventoryhandler.DeleteDuKienTonKho(db)) // DELETE /v1/du-kien-ton-khos/:id: Xóa dự kiến tồn kho
			//dukienTonkhos.GET("/:id", inventoryhandler.GetDuKienTonKho(db))      // GET /v1/du-kien-ton-khos/:id: Lấy chi tiết dự kiến tồn kho
		}
	}

	// Chạy web server ở cổng 8080
	r.Run(":8080")
}
