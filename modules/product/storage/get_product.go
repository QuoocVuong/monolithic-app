package storage

import (
	"context"
	"monolithic-app/modules/product/model"
)

func (s *sqlStore) GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error) {
	var data model.SanPham
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
