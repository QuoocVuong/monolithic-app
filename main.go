package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"monolithic-app/controllers" // Replace with your actual path
	"monolithic-app/models"      // Replace with your actual path
)

func main() {
	// Database setup
	dsn := "root:123456@tcp(localhost:3306)/products_db?charset=utf8mb4&parseTime=True&loc=Local" // Replace with your database credentials
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Trong hàm main(), sau khi khởi tạo productController:
	productController := controllers.NewProductController(db)
	tonKhoController := controllers.NewTonKhoController(db)
	duKienTonKhoController := controllers.NewDuKienTonKhoController(db)
	nhomHangController := controllers.NewNhomHangController(db)

	// Auto Migrate the schema
	db.AutoMigrate(&models.NhomHang{}, &models.MucTuHang{}, &models.SanPham{}, &models.KhoHang{}, &models.TonKho{}, &models.DuKienTonKho{})

	// Gin router setup
	router := gin.Default()

	// Product routes
	router.GET("/products", productController.GetAllProducts)
	router.GET("/products/:id", productController.GetProductByID)
	router.POST("/products", productController.CreateProduct)
	router.PUT("/products/:id", productController.UpdateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)

	// TonKho routes
	router.GET("/tonkho", tonKhoController.GetAllTonKho)
	router.GET("/tonkho/:id", tonKhoController.GetTonKhoByID)
	router.POST("/tonkho", tonKhoController.CreateTonKho)
	router.PUT("/tonkho/:id", tonKhoController.UpdateTonKho)
	router.DELETE("/tonkho/:id", tonKhoController.DeleteTonKho)

	// DuKienTonKho routes
	router.GET("/dukien", duKienTonKhoController.GetAllDuKienTonKho)
	router.GET("/dukien/:id", duKienTonKhoController.GetDuKienTonKhoByID)
	router.POST("/dukien", duKienTonKhoController.CreateDuKienTonKho)
	router.PUT("/dukien/:id", duKienTonKhoController.UpdateDuKienTonKho)
	router.DELETE("/dukien/:id", duKienTonKhoController.DeleteDuKienTonKho)

	// NhomHang routes
	router.GET("/nhomhang", nhomHangController.GetAllNhomHang)
	router.GET("/nhomhang/:id", nhomHangController.GetNhomHangByID)
	router.POST("/nhomhang", nhomHangController.CreateNhomHang)
	router.PUT("/nhomhang/:id", nhomHangController.UpdateNhomHang)
	router.DELETE("/nhomhang/:id", nhomHangController.DeleteNhomHang)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
