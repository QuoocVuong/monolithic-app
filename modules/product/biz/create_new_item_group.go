package biz

import (
	"context"
	"strings"

	"monolithic-app/modules/product/model"
)

type CreateItemGroupStorage interface {
	CreateItemGroup(ctx context.Context, data *model.ItemGroupCreation) error
}
type createItemGroupBiz struct {
	store CreateItemGroupStorage
}

func NewCreateItemGroupBiz(store CreateItemGroupStorage) *createItemGroupBiz {
	return &createItemGroupBiz{store: store}
}
func (biz *createItemGroupBiz) CreateNewItemGroup(ctx context.Context, data *model.ItemGroupCreation) error {
	tennhom := strings.TrimSpace(data.TenNhom)
	if tennhom == "" {
		return model.ErrItemGroupIsBlank
	}
	if err := biz.store.CreateItemGroup(ctx, data); err != nil {
		return err
	}
	return nil
}
