package services

import (
	"gorm.io/gorm"
	"monolithic-app/models" // Thay thế bằng đường dẫn thực tế
)

type NhomHangService struct {
	DB *gorm.DB
}

func NewNhomHangService(db *gorm.DB) *NhomHangService {
	return &NhomHangService{DB: db}
}

// Lấy tất cả nhóm hàng
func (ns *NhomHangService) GetAllNhomHang() ([]models.NhomHang, error) {
	var nhomHang []models.NhomHang
	if err := ns.DB.Find(&nhomHang).Error; err != nil {
		return nil, err
	}
	return nhomHang, nil
}

// Lấy nhóm hàng theo ID
func (ns *NhomHangService) GetNhomHangByID(id uint) (*models.NhomHang, error) {
	var nhomHang models.NhomHang
	if err := ns.DB.First(&nhomHang, id).Error; err != nil {
		return nil, err
	}
	return &nhomHang, nil
}

// Tạo nhóm hàng mới
func (ns *NhomHangService) CreateNhomHang(nhomHang *models.NhomHang) (*models.NhomHang, error) {
	if err := ns.DB.Create(nhomHang).Error; err != nil {
		return nil, err
	}
	return nhomHang, nil
}

// Cập nhật nhóm hàng
func (ns *NhomHangService) UpdateNhomHang(nhomHang *models.NhomHang) error {
	return ns.DB.Save(nhomHang).Error
}

// Xóa nhóm hàng
func (ns *NhomHangService) DeleteNhomHang(id uint) error {
	return ns.DB.Delete(&models.NhomHang{}, id).Error
}
