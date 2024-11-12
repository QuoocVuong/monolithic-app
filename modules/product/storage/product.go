package storage

import (
	"context"
	"gorm.io/gorm"
	"monolithic-app/common"                // Package common chứa các hàm và struct chung
	"monolithic-app/modules/product/model" // Package model chứa các model cho product
)

// sqlStore là struct chứa kết nối database
type sqlStore struct {
	db *gorm.DB
}

// NewSqlStore tạo mới một đối tượng sqlStore với kết nối database
func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

// ======================================= PRODUCT =======================================

// CreateProduct tạo một sản phẩm mới trong database
func (s *sqlStore) CreateProduct(ctx context.Context, data *model.ProductCreation) error {
	// Không cần kiểm tra trùng lặp ma_hang vì database đã có ràng buộc unique

	// Tạo sản phẩm mới bằng GORM
	if err := s.db.Create(data).Error; err != nil {
		return err // Trả về lỗi nếu tạo sản phẩm thất bại
	}

	return nil // Trả về nil nếu tạo sản phẩm thành công
}

// GetProduct lấy thông tin một sản phẩm từ database dựa trên điều kiện cond
func (s *sqlStore) GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error) {
	var data model.SanPham // Biến lưu trữ thông tin sản phẩm

	// Tìm sản phẩm đầu tiên thỏa mãn điều kiện cond bằng GORM
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err // Trả về lỗi nếu tìm kiếm thất bại
	}

	return &data, nil // Trả về sản phẩm và nil nếu tìm kiếm thành công
}

// UpdateProduct cập nhật thông tin một sản phẩm trong database dựa trên điều kiện cond
func (s *sqlStore) UpdateProduct(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ProductUpdate) error {
	// Cập nhật sản phẩm bằng GORM
	if err := s.db.Debug().Where(cond).Updates(dataUpdate).Error; err != nil {
		return err // Trả về lỗi nếu cập nhật thất bại
	}

	return nil // Trả về nil nếu cập nhật thành công
}

// DeleteProduct xóa một sản phẩm trong database dựa trên điều kiện cond (thực hiện soft delete bằng cách cập nhật status)
func (s *sqlStore) DeleteProduct(ctx context.Context, cond map[string]interface{}) error {
	// Cập nhật status của sản phẩm thành "Deleted" bằng GORM
	if err := s.db.Table(model.SanPham{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": model.ProductStatusDeleted}).Error; err != nil {
		return err // Trả về lỗi nếu cập nhật thất bại
	}
	return nil // Trả về nil nếu cập nhật thành công
}

// ListProduct lấy danh sách sản phẩm từ database, hỗ trợ phân trang và lọc
func (s *sqlStore) ListProduct(
	ctx context.Context,
	filter *model.Filterr,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.SanPham, error) {
	var result []model.SanPham

	db := s.db.Table(model.SanPham{}.TableName()) // Sử dụng Table để chỉ định tên bảng rõ ràng

	// --- Lọc sản phẩm chưa bị xóa ---
	//db = db.Where("status != ?", model.ProductStatusDeleted)

	// Áp dụng lọc theo status nếu có (ngoài "deleted")
	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v) // Lọc theo status khác
		}
	}

	// Đếm tổng số sản phẩm (chưa bị xóa) thỏa mãn điều kiện
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	// Lấy danh sách sản phẩm với phân trang
	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Preload("NhomHang").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
