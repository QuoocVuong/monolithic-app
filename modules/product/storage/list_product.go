package storage

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/product/model"
)

func (s *sqlStore) ListProduct(
	ctx context.Context,
	fitler *model.Filterr,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.SanPham, error) {
	var result []model.SanPham

	db := s.db.Where("status <> ?", "Deleted")

	if f := fitler; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.SanPham{}.TableName()).Count(&paging.Total).Error; err != nil {
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
