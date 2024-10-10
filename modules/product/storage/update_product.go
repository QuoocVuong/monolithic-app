package storage

import (
	"context"
	"monolithic-app/modules/product/model"
)

func (s *sqlStore) UpdateProduct(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ProductUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
