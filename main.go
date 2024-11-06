package main

import (
	_ "errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"monolithic-app/configuration"
	"monolithic-app/database"
	"monolithic-app/middleware"
	"monolithic-app/modules/inventory/transport/inventoryhandler"
	"monolithic-app/modules/itemgroup/transport/itemgrouphandler"
	"monolithic-app/modules/product/model"
	"monolithic-app/modules/product/transport/producthandler"
	"time"
)

func main() {
	// 1. Load config từ file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := configuration.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// 2. Kết nối database
	dsn := cfg.GetDSN()
	dbConn, err := database.ConnectToDB(dsn)
	if err != nil {
		panic(err)
	}

	dbConn.AutoMigrate(
		&model.NhomHang{},
		&model.KhoHang{},
		&model.SanPham{},
		&model.TonKho{},
		&model.DuKienTonKho{},
	)

	r := gin.Default()
	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Thay bằng origin của frontend khi deploy
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig)) // CORS middleware trước các middleware khác

	r.Use(middleware.Recover(dbConn))

	v1 := r.Group("/v1")
	v1.Use(middleware.JWTAuthMiddleware())
	{

		products := v1.Group("/products")
		{
			products.POST("", producthandler.CreateProduct(dbConn))
			products.GET("", producthandler.ListProduct(dbConn))
			products.PATCH("/:id", producthandler.UpdateProduct(dbConn))
			products.DELETE("/:id", producthandler.DeleteProduct(dbConn))
			products.GET("/:id", producthandler.GetProduct(dbConn))
		}

		itemgroups := v1.Group("/itemgroups")
		{
			itemgroups.POST("", itemgrouphandler.CreateItemGroup(dbConn))
			itemgroups.GET("", itemgrouphandler.ListItemGroup(dbConn))
			itemgroups.PATCH("/:id", itemgrouphandler.UpdateItemGroup(dbConn))
			itemgroups.DELETE("/:id", itemgrouphandler.DeleteItemGroup(dbConn))
			itemgroups.GET("/:id", itemgrouphandler.GetItemGroup(dbConn))
		}

		khohangs := v1.Group("/khohangs")
		{
			khohangs.POST("", inventoryhandler.CreateKhoHang(dbConn))
			khohangs.GET("", inventoryhandler.ListKhoHang(dbConn))
			khohangs.PATCH("/:id", inventoryhandler.UpdateKhoHang(dbConn))
			khohangs.DELETE("/:id", inventoryhandler.DeleteKhoHang(dbConn))
			khohangs.GET("/:id", inventoryhandler.GetKhoHang(dbConn))
		}

		tonkhos := v1.Group("/tonkhos")
		{
			tonkhos.POST("", inventoryhandler.CreateTonKho(dbConn))
			tonkhos.GET("", inventoryhandler.ListTonKho(dbConn))
			tonkhos.PATCH("/:id", inventoryhandler.UpdateTonKho(dbConn))
			tonkhos.DELETE("/:id", inventoryhandler.DeleteTonKho(dbConn))
			//tonkhos.GET("/:id", inventoryhandler.GetTonKho(dbConn))
		}

		dukienTonkhos := v1.Group("/dukientonkhos")
		{
			dukienTonkhos.POST("", inventoryhandler.CreateDuKienTonKho(dbConn))
			dukienTonkhos.GET("", inventoryhandler.ListDuKienTonKho(dbConn))
			dukienTonkhos.PATCH("/:id", inventoryhandler.UpdateDuKienTonKho(dbConn))
			dukienTonkhos.DELETE("/:id", inventoryhandler.DeleteDuKienTonKho(dbConn))
			//dukienTonkhos.GET("/:id", inventoryhandler.GetDuKienTonKho(dbConn))
		}
	}

	fmt.Println("Server listening on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server:", err)
	}
}
