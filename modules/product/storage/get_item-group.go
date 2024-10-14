package storage

import (
	"context"
	"monolithic-app/modules/product/model"
)

// Implement method GetItemGroup cho *sqlStore
func (s *sqlStore) GetItemGroup(ctx context.Context, cond map[string]interface{}) (*model.NhomHang, error) {
	var data model.NhomHang

	if err := s.db.Table(data.TableName()).Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
