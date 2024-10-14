package biz

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/inventory/model"
)

type TonKhoStorage interface {
	FindTonKho(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.TonKho, error)
	CreateTonKho(ctx context.Context, data *model.TonKho) error
	UpdateTonKho(ctx context.Context, id int, data *model.TonKho) error
	ListTonKho(ctx context.Context, conditions map[string]interface{},
		filter *model.Filterr,
		paging *common.Paging, moreKeys ...string,
	) ([]model.TonKho, error)
	DeleteTonKho(ctx context.Context, id int) error
}

type tonKhoBiz struct {
	store TonKhoStorage
}

func NewTonKhoBiz(store TonKhoStorage) *tonKhoBiz {
	return &tonKhoBiz{store: store}
}

func (biz *tonKhoBiz) CreateNewTonKho(ctx context.Context, data *model.TonKho) error {
	if err := biz.store.CreateTonKho(ctx, data); err != nil {
		return err
	}
	return nil
}

func (biz *tonKhoBiz) UpdateTonKho(ctx context.Context, id int, data *model.TonKho) error {
	if err := biz.store.UpdateTonKho(ctx, id, data); err != nil {
		return err
	}
	return nil
}

func (biz *tonKhoBiz) ListTonKho(ctx context.Context, filter *model.Filterr, paging *common.Paging) ([]model.TonKho, error) {
	data, err := biz.store.ListTonKho(ctx, map[string]interface{}{}, filter, paging, "SanPham", "KhoHang")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (biz *tonKhoBiz) DeleteTonKho(ctx context.Context, id int) error {
	if err := biz.store.DeleteTonKho(ctx, id); err != nil {
		return err
	}
	return nil
}
