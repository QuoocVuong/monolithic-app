package biz

import (
	"context"
	"strings"

	"monolithic-app/common"                // Package common chứa các hàm và struct chung
	"monolithic-app/modules/product/model" // Package model chứa các model cho product
)

// ======================================= PRODUCT =======================================

// ListProductStorage là interface định nghĩa các hàm cần thiết cho việc lấy danh sách sản phẩm.
type ListProductStorage interface {
	ListProduct(
		ctx context.Context,   // Context của request
		filter *model.Filterr, // Biến chứa thông tin lọc
		paging *common.Paging, // Biến chứa thông tin phân trang
		moreKeys ...string,    // Danh sách các trường cần preload
	) ([]model.SanPham, error) // Trả về danh sách sản phẩm và lỗi (nếu có)
}

// listProductBiz là struct chứa logic nghiệp vụ cho việc lấy danh sách sản phẩm.
type listProductBiz struct {
	store ListProductStorage // Interface để truy cập database
}

// NewListProductBiz tạo mới một đối tượng listProductBiz.
func NewListProductBiz(store ListProductStorage) *listProductBiz {
	return &listProductBiz{store: store}
}

// ListProductById là hàm lấy danh sách sản phẩm theo ID.
// Hàm này nhận context, thông tin lọc, và thông tin phân trang làm tham số.
func (biz *listProductBiz) ListProductById(
	ctx context.Context,
	filter *model.Filterr,
	paging *common.Paging,
) ([]model.SanPham, error) {
	// Gọi hàm ListProduct trong store để lấy danh sách sản phẩm từ database
	data, err := biz.store.ListProduct(ctx, filter, paging)
	if err != nil {
		return nil, err // Trả về lỗi nếu lấy danh sách sản phẩm thất bại
	}

	return data, nil // Trả về danh sách sản phẩm và nil nếu thành công
}

// GetProductStorage là interface định nghĩa hàm cần thiết cho việc lấy thông tin một sản phẩm.
type GetProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error)
}

// getProductBiz là struct chứa logic nghiệp vụ cho việc lấy thông tin một sản phẩm.
type getProductBiz struct {
	store GetProductStorage // Interface để truy cập database
}

// NewGetProductBiz tạo mới một đối tượng getProductBiz.
func NewGetProductBiz(store GetProductStorage) *getProductBiz {
	return &getProductBiz{store: store}
}

// GetProductById là hàm lấy thông tin một sản phẩm theo ID.
// Hàm này nhận context và ID làm tham số.
func (biz *getProductBiz) GetProductById(ctx context.Context, id int) (*model.SanPham, error) {
	// Gọi hàm GetProduct trong store để lấy thông tin sản phẩm từ database
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err // Trả về lỗi nếu lấy thông tin sản phẩm thất bại
	}
	return data, nil // Trả về sản phẩm và nil nếu thành công
}

// CreateProductStorage là interface định nghĩa hàm cần thiết cho việc tạo sản phẩm mới.
type CreateProductStorage interface {
	CreateProduct(ctx context.Context, data *model.ProductCreation) error
}

// createProductBiz là struct chứa logic nghiệp vụ cho việc tạo sản phẩm mới.
type createProductBiz struct {
	store CreateProductStorage // Interface để truy cập database
}

// NewCreateProductBiz tạo mới một đối tượng createProductBiz.
func NewCreateProductBiz(store CreateProductStorage) *createProductBiz {
	return &createProductBiz{store: store}
}

// CreateNewProduct là hàm tạo một sản phẩm mới.
// Hàm này nhận context và dữ liệu sản phẩm mới làm tham số.
func (biz *createProductBiz) CreateNewProduct(ctx context.Context, data *model.ProductCreation) error {
	// Kiểm tra mã hàng (ma_hang) có rỗng hay không
	ma_hang := strings.TrimSpace(data.MaHang)
	if ma_hang == "" {
		return model.ErrMaHangIsBlank // Trả về lỗi nếu mã hàng rỗng
	}

	// Gọi hàm CreateProduct trong store để tạo sản phẩm mới trong database
	if err := biz.store.CreateProduct(ctx, data); err != nil {
		return err // Trả về lỗi nếu tạo sản phẩm thất bại
	}

	return nil // Trả về nil nếu tạo sản phẩm thành công
}

// UpdateProductStorage là interface định nghĩa các hàm cần thiết cho việc cập nhật sản phẩm.
type UpdateProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error)
	UpdateProduct(ctx context.Context, cond map[string]interface{}, dataUpdate *model.ProductUpdate) error
}

// updateProductBiz là struct chứa logic nghiệp vụ cho việc cập nhật sản phẩm.
type updateProductBiz struct {
	store UpdateProductStorage // Interface để truy cập database
}

// NewUpdateProductBiz tạo mới một đối tượng updateProductBiz.
func NewUpdateProductBiz(store UpdateProductStorage) *updateProductBiz {
	return &updateProductBiz{store: store}
}

// UpdateProductById là hàm cập nhật thông tin một sản phẩm theo ID.
// Hàm này nhận context, ID, và dữ liệu cập nhật làm tham số.
func (biz *updateProductBiz) UpdateProductById(ctx context.Context, id int, dataUpdate *model.ProductUpdate) error {
	// Lấy thông tin sản phẩm từ database
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err // Trả về lỗi nếu lấy thông tin sản phẩm thất bại
	}

	// Kiểm tra xem sản phẩm đã bị xóa hay chưa
	if data.Status != nil && *data.Status == model.ProductStatusDeleted {
		return model.ErrProductDeleted // Trả về lỗi nếu sản phẩm đã bị xóa
	}

	// Gọi hàm UpdateProduct trong store để cập nhật sản phẩm trong database
	if err := biz.store.UpdateProduct(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err // Trả về lỗi nếu cập nhật sản phẩm thất bại
	}

	return nil // Trả về nil nếu cập nhật sản phẩm thành công
}

// DeleteProductStorage là interface định nghĩa các hàm cần thiết cho việc xóa sản phẩm.
type DeleteProductStorage interface {
	GetProduct(ctx context.Context, cond map[string]interface{}) (*model.SanPham, error)
	DeleteProduct(ctx context.Context, cond map[string]interface{}) error
}

// deleteProductBiz là struct chứa logic nghiệp vụ cho việc xóa sản phẩm.
type deleteProductBiz struct {
	store DeleteProductStorage // Interface để truy cập database
}

// NewDeleteProductBiz tạo mới một đối tượng deleteProductBiz.
func NewDeleteProductBiz(store DeleteProductStorage) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}

// DeleteProductById là hàm xóa một sản phẩm theo ID.
// Hàm này nhận context và ID làm tham số.
func (biz *deleteProductBiz) DeleteProductById(ctx context.Context, id int) error {
	// Lấy thông tin sản phẩm từ database
	data, err := biz.store.GetProduct(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err // Trả về lỗi nếu lấy thông tin sản phẩm thất bại
	}

	// Kiểm tra xem sản phẩm đã bị xóa hay chưa
	if data.Status != nil && *data.Status == model.ProductStatusDeleted {
		return model.ErrProductDeleted // Trả về lỗi nếu sản phẩm đã bị xóa
	}

	// Gọi hàm DeleteProduct trong store để xóa sản phẩm trong database
	if err := biz.store.DeleteProduct(ctx, map[string]interface{}{"id": id}); err != nil {
		return err // Trả về lỗi nếu xóa sản phẩm thất bại
	}

	return nil // Trả về nil nếu xóa sản phẩm thành công
}
