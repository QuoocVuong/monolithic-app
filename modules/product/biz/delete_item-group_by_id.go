package biz

import (
	"context"

	"monolithic-app/modules/product/model"
)

type DeleteItemGroupStorage interface {
	GetItemGroup(ctx context.Context, cond map[string]interface{}) (*model.NhomHang, error)
	DeleteItemGroup(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemGroupBiz struct {
	store DeleteItemGroupStorage
}

func NewDeleteItemGroupBiz(store DeleteItemGroupStorage) *deleteItemGroupBiz {
	return &deleteItemGroupBiz{store: store}
}

func (biz *deleteItemGroupBiz) DeleteItemGroupById(ctx context.Context, id int) error {

	_, err := biz.store.GetItemGroup(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if err := biz.store.DeleteItemGroup(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
