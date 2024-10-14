package storage

import (
	"context"

	"monolithic-app/modules/product/model" // Import đúng package
)

// ... (Các import và code khác)

// Implement method UpdateItemGroup cho *sqlStore
func (s *sqlStore) UpdateItemGroup(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ItemGroupUpdate) error {
	if err := s.db.Table(model.NhomHang{}.TableName()).Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
