package storage

import (
	"context"

	"monolithic-app/common"
	"monolithic-app/modules/product/model" // Đảm bảo import đúng
)

func (s *sqlStore) ListItemGroup(
	ctx context.Context,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.NhomHang, error) {
	var result []model.NhomHang

	db := s.db // Không cần lọc theo status nữa

	if err := db.Table(model.NhomHang{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
