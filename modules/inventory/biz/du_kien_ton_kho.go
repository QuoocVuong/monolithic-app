package biz

import (
	"context"                                // Gói context cho phép truyền thông tin về thời gian timeout, hủy bỏ (cancel), và các giá trị khóa qua nhiều API.
	"monolithic-app/common"                  // Gói chứa các thành phần chung, chẳng hạn như cấu trúc phân trang (paging).
	"monolithic-app/modules/inventory/model" // Gói chứa các mô hình (model) của module inventory (kho hàng), ví dụ cấu trúc DuKienTonKho.
)

// Định nghĩa interface DuKienTonKhoStorage để giao tiếp với tầng lưu trữ (storage layer).
type DuKienTonKhoStorage interface {
	// Tìm kiếm một bản ghi dự kiến tồn kho dựa trên các điều kiện nhất định.
	FindDuKienTonKho(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.DuKienTonKho, error)
	// Tạo mới một bản ghi dự kiến tồn kho.
	CreateDuKienTonKho(ctx context.Context, data *model.DuKienTonKho) error
	// Cập nhật một bản ghi dự kiến tồn kho dựa trên id.
	UpdateDuKienTonKho(ctx context.Context, id int, data *model.DuKienTonKho) error
	// Liệt kê các bản ghi dự kiến tồn kho với điều kiện, bộ lọc, và phân trang.
	ListDuKienTonKho(ctx context.Context, conditions map[string]interface{},
		filter *model.Filterr,
		paging *common.Paging, moreKeys ...string,
	) ([]model.DuKienTonKho, error)
	// Xóa một bản ghi dựa trên id.
	DeleteDuKienTonKho(ctx context.Context, id int) error
}

// Struct chứa logic nghiệp vụ liên quan đến dự kiến tồn kho.
// Nó sử dụng một đối tượng store để giao tiếp với tầng lưu trữ.
type duKienTonKhoBiz struct {
	store DuKienTonKhoStorage // store là một implement của interface DuKienTonKhoStorage.
}

// Constructor tạo mới một đối tượng duKienTonKhoBiz.
// Nó nhận một đối tượng store để tương tác với tầng lưu trữ.
func NewDuKienTonKhoBiz(store DuKienTonKhoStorage) *duKienTonKhoBiz {
	return &duKienTonKhoBiz{store: store} // Khởi tạo và trả về đối tượng duKienTonKhoBiz.
}

// Hàm CreateNewDuKienTonKho dùng để tạo mới một bản ghi dự kiến tồn kho.
// Nó sử dụng store để gọi phương thức CreateDuKienTonKho từ tầng lưu trữ.
func (biz *duKienTonKhoBiz) CreateNewDuKienTonKho(ctx context.Context, data *model.DuKienTonKho) error {
	// Gọi store.CreateDuKienTonKho để lưu bản ghi dự kiến tồn kho vào cơ sở dữ liệu.
	if err := biz.store.CreateDuKienTonKho(ctx, data); err != nil {
		return err // Trả về lỗi nếu có vấn đề trong quá trình tạo.
	}
	return nil // Trả về nil nếu tạo thành công.
}

// Hàm UpdateDuKienTonKho dùng để cập nhật một bản ghi dự kiến tồn kho dựa trên id.
func (biz *duKienTonKhoBiz) UpdateDuKienTonKho(ctx context.Context, id int, data *model.DuKienTonKho) error {
	// Gọi store.UpdateDuKienTonKho để cập nhật bản ghi dựa trên id.
	if err := biz.store.UpdateDuKienTonKho(ctx, id, data); err != nil {
		return err // Trả về lỗi nếu cập nhật thất bại.
	}
	return nil // Trả về nil nếu cập nhật thành công.
}

// Hàm ListDuKienTonKho dùng để liệt kê các bản ghi dự kiến tồn kho với bộ lọc và phân trang.
func (biz *duKienTonKhoBiz) ListDuKienTonKho(ctx context.Context, filter *model.Filterr, paging *common.Paging) ([]model.DuKienTonKho, error) {
	// Gọi store.ListDuKienTonKho để lấy danh sách các bản ghi dựa trên các tham số truyền vào.
	data, err := biz.store.ListDuKienTonKho(ctx, map[string]interface{}{}, filter, paging, "SanPham", "KhoHang")
	if err != nil {
		return nil, err // Trả về lỗi nếu có vấn đề khi lấy dữ liệu.
	}
	return data, nil // Trả về danh sách dữ liệu nếu thành công.
}

// Hàm DeleteDuKienTonKho dùng để xóa một bản ghi dự kiến tồn kho dựa trên id.
func (biz *duKienTonKhoBiz) DeleteDuKienTonKho(ctx context.Context, id int) error {
	// Gọi store.DeleteDuKienTonKho để xóa bản ghi dựa trên id.
	if err := biz.store.DeleteDuKienTonKho(ctx, id); err != nil {
		return err // Trả về lỗi nếu quá trình xóa gặp vấn đề.
	}
	return nil // Trả về nil nếu xóa thành công.
}
