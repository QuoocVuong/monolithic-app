package biz

import (
	"context"
	"monolithic-app/common"
	"monolithic-app/modules/itemgroup/model"
	"strings"
)

// ======================================= ITEM GROUP =======================================

// CreateItemGroupStorage ...
type CreateItemGroupStorage interface {
	CreateItemGroup(ctx context.Context, data *model.ItemGroupCreation) error
}

type createItemGroupBiz struct {
	store CreateItemGroupStorage
}

// NewCreateItemGroupBiz ...
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

// ListItemGroupStorage ...
type ListItemGroupStorage interface {
	ListItemGroup(
		ctx context.Context,
		paging *common.Paging, // Loại bỏ filter
		moreKeys ...string,
	) ([]model.NhomHang, error)
}

type listItemGroupBiz struct {
	store ListItemGroupStorage
}

// NewListItemGroupBiz ...
func NewListItemGroupBiz(store ListItemGroupStorage) *listItemGroupBiz {
	return &listItemGroupBiz{store: store}
}

// Sửa tên hàm và loại bỏ filter
func (biz *listItemGroupBiz) ListItemGroup(
	ctx context.Context,
	paging *common.Paging,
) ([]model.NhomHang, error) { // Sửa kiểu dữ liệu trả về
	data, err := biz.store.ListItemGroup(ctx, paging) // Loại bỏ filter khi gọi store
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetItemGroupStorage ...
type GetItemGroupStorage interface {
	GetItemGroup(ctx context.Context, cond map[string]interface{}) (*model.NhomHang, error)
}

type getItemGroupBiz struct {
	store GetItemGroupStorage
}

// NewGetItemGroupBiz ...
func NewGetItemGroupBiz(store GetItemGroupStorage) *getItemGroupBiz {
	return &getItemGroupBiz{store: store}
}

func (biz *getItemGroupBiz) GetItemGroupById(ctx context.Context, id int) (*model.NhomHang, error) {
	data, err := biz.store.GetItemGroup(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateItemGroupStorage ...
type UpdateItemGroupStorage interface {
	GetItemGroup(ctx context.Context, cond map[string]interface{}) (*model.NhomHang, error)
	UpdateItemGroup(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ItemGroupUpdate) error
}

type updateItemGroupBiz struct {
	store UpdateItemGroupStorage
}

// NewUpdateItemGroupBiz ...
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

// DeleteItemGroupStorage ...
type DeleteItemGroupStorage interface {
	GetItemGroup(ctx context.Context, cond map[string]interface{}) (*model.NhomHang, error)
	DeleteItemGroup(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemGroupBiz struct {
	store DeleteItemGroupStorage
}

// NewDeleteItemGroupBiz ...
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
