package storage

import (
	"context"
	//"gorm.io/gorm"
	"monolithic-app/common"
	"monolithic-app/modules/inventory/model"

	"time"
)

func (s *sqlStore) FindKhoHang(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.KhoHang, error) {
	var data model.KhoHang

	db := s.db.Table(model.KhoHang{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *sqlStore) CreateKhoHang(ctx context.Context, data *model.KhoHangCreate) error {
	db := s.db.Begin()

	if err := db.Table(model.KhoHang{}.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
func (s *sqlStore) UpdateKhoHang(ctx context.Context, id int, data *model.KhoHangUpdate) error {
	db := s.db.Begin()
	if err := db.Table(model.KhoHang{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
func (s *sqlStore) ListKhoHang(
	ctx context.Context,
	filter *model.Filterr,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.KhoHang, error) {
	var result []model.KhoHang
	db := s.db

	if f := filter; f != nil {
		//if v := f.Name; v != "" {
		//	db = db.Where("name = ?", v)
		//}
	}

	if err := db.Table(model.KhoHang{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
func (s *sqlStore) DeleteKhoHang(ctx context.Context, id int) error {
	db := s.db.Begin()
	//if err := db.Table(model.KhoHang{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
	if err := db.Table(model.KhoHang{}.TableName()).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {

		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
