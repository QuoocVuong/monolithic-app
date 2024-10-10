package main

import (
	"fmt"
	"log"
	"monolithic-app/modules/product/model"
	ginproduct "monolithic-app/modules/product/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database setup
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Auto Migrate the schema
	db.AutoMigrate(&model.NhomHang{}, &model.KhoHang{}, &model.SanPham{}, &model.TonKho{}, &model.DuKienTonKho{})
	fmt.Println("Tạo CSDL thành công!")
	//khoi tao gin
	r := gin.Default()
	//tao nhom cho de quan ly
	v1 := r.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", ginproduct.CreateProduct(db))
			products.GET("", ginproduct.ListProduct(db))
			products.PATCH("/:id", ginproduct.UpdateProduct(db))
			products.DELETE("/:id", ginproduct.DeleteProduct(db))
			products.GET("/:id", ginproduct.GetProduct(db))

		}
	}
	r.Run(":8080")
}
