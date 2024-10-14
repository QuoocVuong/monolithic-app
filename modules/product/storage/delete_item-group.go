package storage

import (
	"context"

	"monolithic-app/modules/product/model" // Import đúng package model
)

func (s *sqlStore) DeleteItemGroup(ctx context.Context, cond map[string]interface{}) error {
	// Sử dụng Gorm để xóa nhóm hàng
	if err := s.db.Table(model.NhomHang{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
