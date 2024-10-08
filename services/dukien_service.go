package services

import (
	"gorm.io/gorm"
	"monolithic-app/models" // Thay thế bằng đường dẫn thực tế
)

type DuKienTonKhoService struct {
	DB *gorm.DB
}

func NewDuKienTonKhoService(db *gorm.DB) *DuKienTonKhoService {
	return &DuKienTonKhoService{DB: db}
}

// Lấy tất cả DuKienTonKho
func (ds *DuKienTonKhoService) GetAllDuKienTonKho() ([]models.DuKienTonKho, error) {
	var duKienTonKho []models.DuKienTonKho
	if err := ds.DB.Preload("SanPham").Preload("KhoHang").Find(&duKienTonKho).Error; err != nil {
		return nil, err
	}
	return duKienTonKho, nil
}

// Lấy DuKienTonKho theo ID
func (ds *DuKienTonKhoService) GetDuKienTonKhoByID(id uint) (*models.DuKienTonKho, error) {
	var duKienTonKho models.DuKienTonKho
	if err := ds.DB.Preload("SanPham").Preload("KhoHang").First(&duKienTonKho, id).Error; err != nil {
		return nil, err
	}
	return &duKienTonKho, nil
}

// Tạo DuKienTonKho mới
func (ds *DuKienTonKhoService) CreateDuKienTonKho(duKienTonKho *models.DuKienTonKho) (*models.DuKienTonKho, error) {
	if err := ds.DB.Create(duKienTonKho).Error; err != nil {
		return nil, err
	}
	return duKienTonKho, nil
}

// Cập nhật DuKienTonKho
func (ds *DuKienTonKhoService) UpdateDuKienTonKho(duKienTonKho *models.DuKienTonKho) error {
	return ds.DB.Save(duKienTonKho).Error
}

// Xóa DuKienTonKho
func (ds *DuKienTonKhoService) DeleteDuKienTonKho(id uint) error {
	return ds.DB.Delete(&models.DuKienTonKho{}, id).Error
}
