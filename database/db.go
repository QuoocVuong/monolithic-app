package database

import (
	"log"

	"gorm.io/driver/mysql" // Driver MySQL cho GORM
	"gorm.io/gorm"         // GORM ORM cho database
)

// ConnectToDB kết nối đến database MySQL dựa trên DSN đã cho và trả về đối tượng *gorm.DB.
func ConnectToDB(dsn string) *gorm.DB {
	// Mở kết nối đến database MySQL sử dụng gorm.Open với driver MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err) // Nếu lỗi kết nối, in lỗi và dừng chương trình
	}

	return db // Trả về đối tượng *gorm.DB đại diện cho kết nối
}
