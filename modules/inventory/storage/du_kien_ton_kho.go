package storage

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/inventory/model"
)

func (s *sqlStore) FindDuKienTonKho(
	ctx context.Context, conditions map[string]interface{}, moreKeys ...string,
) (*model.DuKienTonKho, error) {
	var data model.DuKienTonKho
	db := s.db.Table(model.DuKienTonKho{}.TableName())
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *sqlStore) CreateDuKienTonKho(
	ctx context.Context, data *model.DuKienTonKho,
) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}

func (s *sqlStore) UpdateDuKienTonKho(
	ctx context.Context,
	id int, data *model.DuKienTonKho,
) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}

func (s *sqlStore) ListDuKienTonKho(
	ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filterr,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.DuKienTonKho, error) {
	var result []model.DuKienTonKho
	db := s.db.Table(model.DuKienTonKho{}.TableName())
	db = db.Where(conditions)

	if f := filter; f != nil {
		//if v := f.Name; v != "" {
		//	db = db.Where("name = ?", v)
		//}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.Sort; v != "" {
		if err := db.Order(v).Error; err != nil {
			return nil, err
		}
	} else {
		db.Order("id desc")
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sqlStore) DeleteDuKienTonKho(
	ctx context.Context, id int,
) error {
	db := s.db.Begin()

	if err := db.Table(model.DuKienTonKho{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
