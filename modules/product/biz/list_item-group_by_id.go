package biz

import (
	"context"

	"monolithic-app/common"
	"monolithic-app/modules/product/model" // Sửa import cho đúng package
)

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
