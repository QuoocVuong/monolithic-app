package biz

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/inventory/model"
)

type KhoHangStorage interface {
	FindKhoHang(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.KhoHang, error)
	CreateKhoHang(ctx context.Context, data *model.KhoHangCreate) error
	UpdateKhoHang(ctx context.Context, id int, data *model.KhoHangUpdate) error
	ListKhoHang(
		ctx context.Context,
		filter *model.Filterr,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.KhoHang, error)
	DeleteKhoHang(ctx context.Context, id int) error
}

type khoHangBiz struct {
	store KhoHangStorage
}

func NewKhoHangBiz(store KhoHangStorage) *khoHangBiz {
	return &khoHangBiz{store: store}
}

func (biz *khoHangBiz) CreateNewKhoHang(ctx context.Context, data *model.KhoHangCreate) error {
	// Kiểm tra xem tên kho hàng đã tồn tại hay chưa
	if _, err := biz.store.FindKhoHang(ctx, map[string]interface{}{"ten_kho": data.TenKho}); err == nil {
		return model.ErrKhoHangExisted
	}

	if err := biz.store.CreateKhoHang(ctx, data); err != nil {
		return err
	}

	return nil
}
func (biz *khoHangBiz) UpdateKhoHang(ctx context.Context, id int, data *model.KhoHangUpdate) error {
	if err := biz.store.UpdateKhoHang(ctx, id, data); err != nil {
		return err
	}
	return nil
}
func (biz *khoHangBiz) ListKhoHang(ctx context.Context,
	filter *model.Filterr,
	paging *common.Paging) ([]model.KhoHang, error) {
	data, err := biz.store.ListKhoHang(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (biz *khoHangBiz) DeleteKhoHang(ctx context.Context, id int) error {
	if err := biz.store.DeleteKhoHang(ctx, id); err != nil {
		return err
	}
	return nil
}
