package storage

import (
	"context"
	"monolithic-app/modules/product/model"
)

func (s *sqlStore) CreateProduct(ctx context.Context, data *model.ProductCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
