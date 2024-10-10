package biz

import (
	"context"
	"monolithic-app/modules/product/model"
)

type GetProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error)
}
type getProductBiz struct {
	store GetProductStorage
}

func NewGetProductBiz(store GetProductStorage) *getProductBiz {

	return &getProductBiz{store: store}
}

func (biz *getProductBiz) GetProductById(ctx context.Context, id int) (*model.SanPham, error) {

	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
