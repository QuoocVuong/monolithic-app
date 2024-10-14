package storage

import (
	"context"

	"monolithic-app/common"                  // Package common chứa các struct và hàm chung
	"monolithic-app/modules/itemgroup/model" // Package model chứa các model cho itemgroup
)

// ======================================= ITEM GROUP =======================================

// CreateItemGroup tạo mới một nhóm hàng trong database.
func (s *sqlStore) CreateItemGroup(ctx context.Context, data *model.ItemGroupCreation) error {
	// Sử dụng GORM để tạo một bản ghi mới trong database
	if err := s.db.Create(&data).Error; err != nil {
		return err // Trả về lỗi nếu tạo nhóm hàng thất bại
	}

	return nil // Trả về nil nếu tạo nhóm hàng thành công
}

// GetItemGroup lấy thông tin một nhóm hàng từ database dựa trên điều kiện cond.
func (s *sqlStore) GetItemGroup(ctx context.Context, cond map[string]interface{}) (*model.NhomHang, error) {
	var data model.NhomHang // Biến để lưu trữ thông tin nhóm hàng

	// Sử dụng GORM để tìm kiếm nhóm hàng đầu tiên thỏa mãn điều kiện cond
	if err := s.db.Table(data.TableName()).Where(cond).First(&data).Error; err != nil {
		return nil, err // Trả về lỗi nếu tìm kiếm thất bại
	}

	return &data, nil // Trả về nhóm hàng và nil nếu tìm kiếm thành công
}

// UpdateItemGroup cập nhật thông tin một nhóm hàng trong database dựa trên điều kiện cond.
func (s *sqlStore) UpdateItemGroup(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ItemGroupUpdate) error {
	// Sử dụng GORM để cập nhật nhóm hàng thỏa mãn điều kiện cond
	if err := s.db.Table(model.NhomHang{}.TableName()).Where(cond).Updates(dataUpdate).Error; err != nil {
		return err // Trả về lỗi nếu cập nhật thất bại
	}

	return nil // Trả về nil nếu cập nhật thành công
}

// DeleteItemGroup xóa một nhóm hàng trong database dựa trên điều kiện cond.
func (s *sqlStore) DeleteItemGroup(ctx context.Context, cond map[string]interface{}) error {
	// Sử dụng GORM để xóa nhóm hàng thỏa mãn điều kiện cond
	if err := s.db.Table(model.NhomHang{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return err // Trả về lỗi nếu xóa thất bại
	}

	return nil // Trả về nil nếu xóa thành công
}

// ListItemGroup lấy danh sách nhóm hàng từ database, hỗ trợ phân trang.
func (s *sqlStore) ListItemGroup(
	ctx context.Context,
	paging *common.Paging, // Biến chứa thông tin phân trang
	moreKeys ...string, // Danh sách các trường cần preload
) ([]model.NhomHang, error) {
	var result []model.NhomHang // Biến để lưu trữ danh sách nhóm hàng

	db := s.db // Sử dụng đối tượng database từ sqlStore

	// Đếm tổng số nhóm hàng trong database
	if err := db.Table(model.NhomHang{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err // Trả về lỗi nếu đếm thất bại
	}

	// Sử dụng GORM để lấy danh sách nhóm hàng với phân trang
	if err := db.Order("id desc"). // Sắp xếp theo ID giảm dần
					Offset((paging.Page - 1) * paging.Limit). // Bỏ qua số lượng bản ghi bằng (page-1)*limit
					Limit(paging.Limit).                      // Giới hạn số lượng bản ghi trả về bằng limit
					Find(&result).Error; err != nil {         // Tìm kiếm danh sách nhóm hàng
		return nil, err // Trả về lỗi nếu tìm kiếm thất bại
	}

	return result, nil // Trả về danh sách nhóm hàng và nil nếu tìm kiếm thành công
}
