package services

import (
	"gorm.io/gorm"
	"monolithic-app/models" // Thay thế bằng đường dẫn thực tế
)

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

// Lấy tất cả sản phẩm
func (ps *ProductService) GetAllProducts() ([]models.SanPham, error) {
	var products []models.SanPham
	if err := ps.DB.Preload("NhomHang").Preload("MucTuHang").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Lấy sản phẩm theo ID
func (ps *ProductService) GetProductByID(id uint) (*models.SanPham, error) {
	var product models.SanPham
	if err := ps.DB.Preload("NhomHang").Preload("MucTuHang").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Tạo sản phẩm mới
func (ps *ProductService) CreateProduct(product *models.SanPham) (*models.SanPham, error) {
	if err := ps.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// Cập nhật sản phẩm
func (ps *ProductService) UpdateProduct(product *models.SanPham) error {
	return ps.DB.Save(product).Error
}

// Xóa sản phẩm
func (ps *ProductService) DeleteProduct(id uint) error {
	return ps.DB.Delete(&models.SanPham{}, id).Error
}
