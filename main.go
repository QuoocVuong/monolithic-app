package main

import (
	"log"
	"monolithic-app/modules/inventory/transport/inventoryhandler"

	"monolithic-app/modules/product/model"
	ginproduct "monolithic-app/modules/product/transport/producthandler"

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
	db.AutoMigrate(
		&model.NhomHang{},
		&model.KhoHang{},
		&model.SanPham{},
		&model.TonKho{},
		&model.DuKienTonKho{})

	//khoi tao producthandler
	r := gin.Default()
	//tao nhom cho de quan ly

	v1 := r.Group("/v1")
	{
		// Routes cho product
		products := v1.Group("/products")
		{
			products.POST("", ginproduct.CreateProduct(db))
			products.GET("", ginproduct.ListProduct(db))
			products.PATCH("/:id", ginproduct.UpdateProduct(db))
			products.DELETE("/:id", ginproduct.DeleteProduct(db))
			products.GET("/:id", ginproduct.GetProduct(db))

		}
		// Routes cho itemgroup
		itemgroups := v1.Group("/itemgroups")
		{
			itemgroups.POST("", ginproduct.CreateItemGroup(db))
			itemgroups.GET("", ginproduct.ListItemGroup(db))
			itemgroups.PATCH("/:id", ginproduct.UpdateItemGroup(db))
			itemgroups.DELETE("/:id", ginproduct.DeleteItemGroup(db))
			itemgroups.GET("/:id")

		}
		// Routes cho kho_hang
		khohangs := v1.Group("/kho-hangs")
		{
			khohangs.POST("", inventoryhandler.CreateKhoHang(db))
			khohangs.GET("", inventoryhandler.ListKhoHang(db))
			khohangs.PATCH("/:id", inventoryhandler.UpdateKhoHang(db))
			khohangs.DELETE("/:id", inventoryhandler.DeleteKhoHang(db))
			//khohangs.GET("/:id", inventoryhandler.GetKhoHang(db)) // Nếu cần
		}

		// Routes cho ton_kho
		tonkhos := v1.Group("/ton-khos")
		{
			tonkhos.POST("", inventoryhandler.CreateTonKho(db))
			tonkhos.GET("", inventoryhandler.ListTonKho(db))
			tonkhos.PATCH("/:id", inventoryhandler.UpdateTonKho(db))
			tonkhos.DELETE("/:id", inventoryhandler.DeleteTonKho(db))
			//tonkhos.GET("/:id", inventoryhandler.GetTonKho(db)) // Nếu cần
		}

		// Routes cho du_kien_ton_kho
		dukienTonkhos := v1.Group("/du-kien-ton-khos")
		{
			dukienTonkhos.POST("", inventoryhandler.CreateDuKienTonKho(db))
			dukienTonkhos.GET("", inventoryhandler.ListDuKienTonKho(db))
			dukienTonkhos.PATCH("/:id", inventoryhandler.UpdateDuKienTonKho(db))
			dukienTonkhos.DELETE("/:id", inventoryhandler.DeleteDuKienTonKho(db))
			//dukienTonkhos.GET("/:id", inventoryhandler.GetDuKienTonKho(db)) // Nếu cần
		}
	}
	r.Run(":8080")
}
