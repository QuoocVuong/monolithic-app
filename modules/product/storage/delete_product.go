package storage

import (
	"context"
	"monolithic-app/modules/product/model"
)

func (s *sqlStore) DeleteProduct(ctx context.Context, cond map[string]interface{}) error {

	if err := s.db.Table(model.SanPham{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": model.ProductStatusDeleted}).Error; err != nil {
		return err
	}
	return nil
}
