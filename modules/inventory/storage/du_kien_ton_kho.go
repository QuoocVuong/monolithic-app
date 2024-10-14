package storage

import (
	"context"

	"monolithic-app/common"                  // Package common chứa các struct và hàm chung
	"monolithic-app/modules/inventory/model" // Package model chứa các model cho inventory
)

// ======================================= DỰ KIẾN TỒN KHO =======================================

// FindDuKienTonKho tìm kiếm một dự kiến tồn kho theo điều kiện `conditions` và preload các trường liên quan `moreKeys`.
func (s *sqlStore) FindDuKienTonKho(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.DuKienTonKho, error) {
	var data model.DuKienTonKho // Biến để lưu trữ dự kiến tồn kho

	db := s.db.Table(model.DuKienTonKho{}.TableName()) // Lấy table dự kiến tồn kho

	// Preload các trường liên quan (nếu có)
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	// Tìm kiếm dự kiến tồn kho đầu tiên thỏa mãn điều kiện
	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, err // Trả về lỗi nếu tìm kiếm thất bại
	}

	return &data, nil // Trả về dự kiến tồn kho và nil nếu tìm kiếm thành công
}

// CreateDuKienTonKho tạo mới một dự kiến tồn kho.
func (s *sqlStore) CreateDuKienTonKho(ctx context.Context, data *model.DuKienTonKho) error {
	db := s.db.Begin() // Bắt đầu transaction

	// Tạo dự kiến tồn kho mới
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback() // Rollback transaction nếu có lỗi
		return err    // Trả về lỗi
	}

	// Commit transaction
	if err := db.Commit().Error; err != nil {
		db.Rollback() // Rollback transaction nếu có lỗi
		return err    // Trả về lỗi
	}

	return nil // Trả về nil nếu tạo thành công
}

// UpdateDuKienTonKho cập nhật thông tin của một dự kiến tồn kho theo ID.
func (s *sqlStore) UpdateDuKienTonKho(ctx context.Context, id int, data *model.DuKienTonKho) error {
	db := s.db.Begin() // Bắt đầu transaction

	// Cập nhật dự kiến tồn kho
	if err := db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		db.Rollback() // Rollback transaction nếu có lỗi
		return err    // Trả về lỗi
	}

	// Commit transaction
	if err := db.Commit().Error; err != nil {
		db.Rollback() // Rollback transaction nếu có lỗi
		return err    // Trả về lỗi
	}

	return nil // Trả về nil nếu cập nhật thành công
}

// ListDuKienTonKho lấy danh sách dự kiến tồn kho, hỗ trợ phân trang và lọc.
func (s *sqlStore) ListDuKienTonKho(
	ctx context.Context,
	conditions map[string]interface{}, // Điều kiện lọc
	filter *model.Filterr, // Biến chứa thông tin lọc (chưa được sử dụng)
	paging *common.Paging, // Biến chứa thông tin phân trang
	moreKeys ...string, // Danh sách các trường cần preload
) ([]model.DuKienTonKho, error) {
	var result []model.DuKienTonKho                    // Biến để lưu trữ danh sách dự kiến tồn kho
	db := s.db.Table(model.DuKienTonKho{}.TableName()) // Lấy table dự kiến tồn kho
	db = db.Where(conditions)                          // Áp dụng điều kiện lọc (nếu có)

	// // Áp dụng lọc (chưa được sử dụng)
	// if f := filter; f != nil {
	// 	//if v := f.Name; v != "" {
	// 	//	db = db.Where("name = ?", v)
	// 	//}
	// }

	// Đếm tổng số dự kiến tồn kho
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err // Trả về lỗi nếu đếm thất bại
	}

	// Preload các trường liên quan (nếu có)
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	// Áp dụng sắp xếp
	if v := paging.Sort; v != "" {
		if err := db.Order(v).Error; err != nil { // Sắp xếp theo trường được chỉ định
			return nil, err // Trả về lỗi nếu sắp xếp thất bại
		}
	} else {
		db.Order("id desc") // Sắp xếp theo ID giảm dần nếu không có trường sắp xếp được chỉ định
	}

	// Lấy danh sách dự kiến tồn kho với phân trang
	if err := db.
		Offset((paging.Page - 1) * paging.Limit). // Bỏ qua (page-1)*limit bản ghi đầu tiên
		Limit(paging.Limit).                      // Giới hạn số lượng bản ghi trả về bằng limit
		Find(&result).Error; err != nil {         // Tìm kiếm danh sách dự kiến tồn kho
		return nil, err // Trả về lỗi nếu tìm kiếm thất bại
	}

	return result, nil // Trả về danh sách dự kiến tồn kho và nil nếu thành công
}

// DeleteDuKienTonKho xóa một dự kiến tồn kho theo ID.
func (s *sqlStore) DeleteDuKienTonKho(ctx context.Context, id int) error {
	db := s.db.Begin() // Bắt đầu transaction

	// Xóa dự kiến tồn kho
	if err := db.Table(model.DuKienTonKho{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		db.Rollback() // Rollback transaction nếu có lỗi
		return err    // Trả về lỗi
	}

	// Commit transaction
	if err := db.Commit().Error; err != nil {
		db.Rollback() // Rollback transaction nếu có lỗi
		return err    // Trả về lỗi
	}

	return nil // Trả về nil nếu xóa thành công
}
