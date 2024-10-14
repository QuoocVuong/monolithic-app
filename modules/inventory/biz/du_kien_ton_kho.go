package biz

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/inventory/model"
)

type DuKienTonKhoStorage interface {
	FindDuKienTonKho(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.DuKienTonKho, error)
	CreateDuKienTonKho(ctx context.Context, data *model.DuKienTonKho) error
	UpdateDuKienTonKho(ctx context.Context, id int, data *model.DuKienTonKho) error
	ListDuKienTonKho(ctx context.Context, conditions map[string]interface{},
		filter *model.Filterr,
		paging *common.Paging, moreKeys ...string,
	) ([]model.DuKienTonKho, error)
	DeleteDuKienTonKho(ctx context.Context, id int) error
}

type duKienTonKhoBiz struct {
	store DuKienTonKhoStorage
}

func NewDuKienTonKhoBiz(store DuKienTonKhoStorage) *duKienTonKhoBiz {
	return &duKienTonKhoBiz{store: store}
}

func (biz *duKienTonKhoBiz) CreateNewDuKienTonKho(ctx context.Context, data *model.DuKienTonKho) error {
	if err := biz.store.CreateDuKienTonKho(ctx, data); err != nil {
		return err
	}
	return nil
}
func (biz *duKienTonKhoBiz) UpdateDuKienTonKho(ctx context.Context, id int, data *model.DuKienTonKho) error {
	if err := biz.store.UpdateDuKienTonKho(ctx, id, data); err != nil {
		return err
	}
	return nil
}
func (biz *duKienTonKhoBiz) ListDuKienTonKho(ctx context.Context, filter *model.Filterr, paging *common.Paging) ([]model.DuKienTonKho, error) {
	data, err := biz.store.ListDuKienTonKho(ctx, map[string]interface{}{}, filter, paging, "SanPham", "KhoHang")
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (biz *duKienTonKhoBiz) DeleteDuKienTonKho(ctx context.Context, id int) error {
	if err := biz.store.DeleteDuKienTonKho(ctx, id); err != nil {
		return err
	}
	return nil
}
