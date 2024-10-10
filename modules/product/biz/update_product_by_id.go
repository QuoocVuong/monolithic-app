package biz

import (
	"context"
	"monolithic-app/modules/product/model"
)

type UpdateProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error)
	UpdateProduct(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ProductUpdate) error
}
type updateProductBiz struct {
	store UpdateProductStorage
}

func NewUpdateProductBiz(store UpdateProductStorage) *updateProductBiz {

	return &updateProductBiz{store: store}
}
func (biz *updateProductBiz) UpdateProductById(ctx context.Context, id int, dataUpdate *model.ProductUpdate) error {
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ProductStatusDeleted {
		return model.ErrProductDeleted
	}
	if err := biz.store.UpdateProduct(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
