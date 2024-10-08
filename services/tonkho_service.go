package services

import (
	"gorm.io/gorm"
	"monolithic-app/models" // Thay thế bằng đường dẫn thực tế
)

type TonKhoService struct {
	DB *gorm.DB
}

func NewTonKhoService(db *gorm.DB) *TonKhoService {
	return &TonKhoService{DB: db}
}

// Lấy tất cả TonKho
func (ts *TonKhoService) GetAllTonKho() ([]models.TonKho, error) {
	var tonKho []models.TonKho
	if err := ts.DB.Preload("SanPham").Preload("KhoHang").Find(&tonKho).Error; err != nil {
		return nil, err
	}
	return tonKho, nil
}

// Lấy TonKho theo ID
func (ts *TonKhoService) GetTonKhoByID(id uint) (*models.TonKho, error) {
	var tonKho models.TonKho
	if err := ts.DB.Preload("SanPham").Preload("KhoHang").First(&tonKho, id).Error; err != nil {
		return nil, err
	}
	return &tonKho, nil
}

// Tạo TonKho mới
func (ts *TonKhoService) CreateTonKho(tonKho *models.TonKho) (*models.TonKho, error) {
	if err := ts.DB.Create(tonKho).Error; err != nil {
		return nil, err
	}
	return tonKho, nil
}

// Cập nhật TonKho
func (ts *TonKhoService) UpdateTonKho(tonKho *models.TonKho) error {
	return ts.DB.Save(tonKho).Error
}

// Xóa TonKho
func (ts *TonKhoService) DeleteTonKho(id uint) error {
	return ts.DB.Delete(&models.TonKho{}, id).Error
}
