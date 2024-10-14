package storage

import "gorm.io/gorm" // GORM ORM cho database

// sqlStore là struct chứa kết nối database
type sqlStore struct {
	db *gorm.DB
}

// NewSqlStore tạo mới một đối tượng sqlStore với kết nối database
func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
