package storage

import (
	"context"

	"monolithic-app/modules/product/model"
)

func (s *sqlStore) CreateProduct(ctx context.Context, data *model.ProductCreation) error {
	// Không cần kiểm tra trùng lặp ma_hang

	// Tạo sản phẩm mới
	if err := s.db.Create(data).Error; err != nil {
		return err // Trả về lỗi database nếu có
	}

	return nil // Trả về nil nếu tạo thành công
}
