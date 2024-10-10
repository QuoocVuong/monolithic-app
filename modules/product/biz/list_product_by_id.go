package biz

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/product/model"
)

type ListProductStorage interface {
	ListProduct(
		ctx context.Context,
		fitler *model.Filterr,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.SanPham, error)
}
type listProductBiz struct {
	store ListProductStorage
}

func NewListProductBiz(store ListProductStorage) *listProductBiz {

	return &listProductBiz{store: store}
}

func (biz *listProductBiz) ListProductById(
	ctx context.Context,
	fitler *model.Filterr,
	paging *common.Paging,
) ([]model.SanPham, error) {
	data, err := biz.store.ListProduct(ctx, fitler, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
