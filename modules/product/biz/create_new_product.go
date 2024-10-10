package biz

import (
	"context"
	"monolithic-app/modules/product/model"
	"strings"
)

type CreateProductStorage interface {
	CreateProduct(ctx context.Context, data *model.ProductCreation) error
}
type createProductBiz struct {
	store CreateProductStorage
}

func CreateNewProductBiz(store CreateProductStorage) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateNewProduct(ctx context.Context, data *model.ProductCreation) error {
	mahang := strings.TrimSpace(data.MaHang)
	if mahang == "" {
		return model.ErrMaHangIsBlank
	}
	if err := biz.store.CreateProduct(ctx, data); err != nil {
		return err
	}
	return nil
}
