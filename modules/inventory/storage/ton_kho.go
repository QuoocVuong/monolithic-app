package storage

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/inventory/model"
)

func (s *sqlStore) FindTonKho(
	ctx context.Context, conditions map[string]interface{}, moreKeys ...string,
) (*model.TonKho, error) {
	var data model.TonKho
	db := s.db.Table(model.TonKho{}.TableName())
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *sqlStore) CreateTonKho(
	ctx context.Context, data *model.TonKho,
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

func (s *sqlStore) UpdateTonKho(
	ctx context.Context,
	id int, data *model.TonKho,
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

func (s *sqlStore) ListTonKho(
	ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filterr,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TonKho, error) {
	var result []model.TonKho

	db := s.db.Table(model.TonKho{}.TableName())
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
		if err := db.Order(v); err != nil {
			return nil, err
		}
	} else {
		db = db.Order("id desc")
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sqlStore) DeleteTonKho(
	ctx context.Context, id int,
) error {
	db := s.db.Begin()

	if err := db.Table(model.TonKho{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
