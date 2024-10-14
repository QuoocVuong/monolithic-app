package biz

import (
	"context"

	"monolithic-app/modules/product/model"
)

type UpdateItemGroupStorage interface {
	GetItemGroup(ctx context.Context, cond map[string]interface{}) (*model.NhomHang, error)
	UpdateItemGroup(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ItemGroupUpdate) error
}

type updateItemGroupBiz struct {
	store UpdateItemGroupStorage
}

func NewUpdateItemGroupBiz(store UpdateItemGroupStorage) *updateItemGroupBiz {
	return &updateItemGroupBiz{store: store}
}

func (biz *updateItemGroupBiz) UpdateItemGroupById(ctx context.Context, id int, dataUpdate *model.ItemGroupUpdate) error {
	_, err := biz.store.GetItemGroup(ctx, map[string]interface{}{"id": id}) // Không cần gán data
	if err != nil {
		return err
	}

	// Loại bỏ phần kiểm tra status:
	// if data.Status != nil && *data.Status == model.ProductStatusDeleted { ... }

	if err := biz.store.UpdateItemGroup(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
