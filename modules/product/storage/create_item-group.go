package storage

import (
	"context"
	"monolithic-app/modules/product/model"
)

//	func (s *sqlStore) CreateItemGroup(ctx context.Context, data *model.ItemGroupCreation) error {
//		if err := s.db.Create(&data).Error; err != nil {
//			return err
//		}
//		return nil
//	}
func (s *sqlStore) CreateItemGroup(ctx context.Context, data *model.ItemGroupCreation) error {
	// Sử dụng Gorm để tạo record mới
	if err := s.db.Create(&data).Error; err != nil {

		return err
	}

	return nil
}
