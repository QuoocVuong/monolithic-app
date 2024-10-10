package biz

import (
	"context"
	"monolithic-app/modules/product/model"
)

type DeleteProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error)
	DeleteProduct(ctx context.Context, cond map[string]interface{}) error
}
type deleteProductBiz struct {
	store DeleteProductStorage
}

func NewDeleteProductBiz(store DeleteProductStorage) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}

func (biz *deleteProductBiz) DeleteProductById(ctx context.Context, id int) error {
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ProductStatusDeleted {
		return model.ErrProductDeleted
	}
	if err := biz.store.DeleteProduct(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
