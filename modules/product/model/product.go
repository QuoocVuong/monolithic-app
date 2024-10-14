package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"monolithic-app/common" // Package common chứa các struct và hàm chung

	"monolithic-app/modules/itemgroup/model" // Import model từ module itemgroup
)

// ======================================= PRODUCT =======================================

// ErrMaHangIsBlank là lỗi khi mã hàng bị bỏ trống
var (
	ErrMaHangIsBlank  = errors.New("MaHang CanNot Is Blank")
	ErrProductDeleted = errors.New("product is deleted")
)

// NhomHang là struct đại diện cho bảng "nhom_hangs" trong database
type NhomHang struct {
	common.SQLModel        // Struct nhúng từ common, chứa các trường ID, CreatedAt, UpdatedAt, DeletedAt
	TenNhom         string `gorm:"column:ten_nhom;type:varchar(255)" json:"tenNhom"` // Tên
}

// SanPham là struct đại diện cho bảng "san_phams" trong database
type SanPham struct {
	common.SQLModel                      // Struct nhúng từ common, chứa các trường ID, CreatedAt, UpdatedAt, DeletedAt
	MaHang                string         `gorm:"column:ma_hang;type:varchar(255)" json:"maHang"`                                             // Mã hàng
	NhomHangID            uint           `gorm:"column:nhom_hang_id;index;foreignKey:NhomHangID" json:"nhomHangID"`                          // ID của nhóm hàng
	TenSanPham            string         `gorm:"column:ten_san_pham;type:varchar(255)" json:"tenSanPham"`                                    // Tên sản phẩm
	DonViTinh             string         `gorm:"column:don_vi_tinh;type:varchar(255)" json:"donViTinh"`                                      // Đơn vị tính
	CoPhieuMoDau          float64        `gorm:"column:co_phieu_mo_dau" json:"coPhieuMoDau"`                                                 // Cổ phiếu mở đầu
	DinhGia               float64        `gorm:"column:dinh_gia" json:"dinhGia"`                                                             // Định giá
	TyGiaBanHangTieuChuan float64        `gorm:"column:ty_gia_ban_hang_tieu_chuan" json:"tyGiaBanHangTieuChuan"`                             // Tỷ giá bán hàng tiêu chuẩn
	ChiDinhLoaiTaiSan     string         `gorm:"column:chi_dinh_loai_tai_san;type:varchar(255)" json:"ChiDinhLoaiTaiSan"`                    // Chỉ định loại tài sản
	Status                *ProductStatus `gorm:"column:status;type:enum('selling','out_of_stock','deleted');default:'selling" json:"status"` // Trạng thái sản phẩm
	HanSuDung             *time.Time     `gorm:"column:han_su_dung" json:"hanSuDung"`                                                        // Hạn sử dụng
	NhomHang              model.NhomHang `gorm:"foreignKey:NhomHangID" json:"nhomHang"`                                                      // Nhóm hàng (foreign key)
}

// TableName định nghĩa tên bảng cho model SanPham
func (SanPham) TableName() string { return "san_phams" }

// ProductCreation là struct chứa dữ liệu để tạo mới sản phẩm
type ProductCreation struct {
	common.SQLModel                      // Struct nhúng từ common, chứa các trường ID, CreatedAt, UpdatedAt, DeletedAt
	MaHang                string         `gorm:"column:ma_hang;type:varchar(255)" json:"maHang"`                                             // Mã hàng
	NhomHangID            uint           `gorm:"column:nhom_hang_id;index;foreignKey:NhomHangID" json:"nhomHangID"`                          // ID của nhóm hàng
	TenSanPham            string         `gorm:"column:ten_san_pham;type:varchar(255)" json:"tenSanPham"`                                    // Tên sản phẩm
	DonViTinh             string         `gorm:"column:don_vi_tinh;type:varchar(255)" json:"donViTinh"`                                      // Đơn vị tính
	CoPhieuMoDau          float64        `gorm:"column:co_phieu_mo_dau" json:"coPhieuMoDau"`                                                 // Cổ phiếu mở đầu
	DinhGia               float64        `gorm:"column:dinh_gia" json:"dinhGia"`                                                             // Định giá
	TyGiaBanHangTieuChuan float64        `gorm:"column:ty_gia_ban_hang_tieu_chuan" json:"tyGiaBanHangTieuChuan"`                             // Tỷ giá bán hàng tiêu chuẩn
	ChiDinhLoaiTaiSan     string         `gorm:"column:chi_dinh_loai_tai_san;type:varchar(255)" json:"ChiDinhLoaiTaiSan"`                    // Chỉ định loại tài sản
	HanSuDung             *time.Time     `gorm:"column:han_su_dung" json:"hanSuDung"`                                                        // Hạn sử dụng
	Status                *ProductStatus `gorm:"column:status;type:enum('selling','out_of_stock','deleted');default:'selling" json:"status"` // Trạng thái sản phẩm
}

// TableName định nghĩa tên bảng cho model ProductCreation
func (ProductCreation) TableName() string { return SanPham{}.TableName() }

// ProductUpdate là struct chứa dữ liệu để cập nhật sản phẩm
// Sử dụng con trỏ cho phép cập nhật các trường thành giá trị rỗng
type ProductUpdate struct {
	common.SQLModel                      // Struct nhúng từ common, chứa các trường ID, CreatedAt, UpdatedAt, DeletedAt
	MaHang                string         `gorm:"column:ma_hang " json:"maHang"`                                                              // Mã hàng
	NhomHangID            uint           `gorm:"column:nhom_hang_id;index;foreignKey:NhomHangID" json:"nhomHangID"`                          // ID của nhóm hàng
	TenSanPham            string         `gorm:"column:ten_san_pham " json:"tenSanPham"`                                                     // Tên sản phẩm
	DonViTinh             string         `gorm:"column:don_vi_tinh " json:"donViTinh"`                                                       // Đơn vị tính
	CoPhieuMoDau          string         `gorm:"column:co_phieu_mo_dau " json:"coPhieuMoDau"`                                                // Cổ phiếu mở đầu
	DinhGia               string         `gorm:"column:dinh_gia " json:"dinhGia"`                                                            // Định giá
	TyGiaBanHangTieuChuan float64        `gorm:"column:ty_gia_ban_hang_tieu_chuan" json:"tyGiaBanHangTieuChuan"`                             // Tỷ giá bán hàng tiêu chuẩn
	ChiDinhLoaiTaiSan     string         `gorm:"column:chi_dinh_loai_tai_san " json:"ChiDinhLoaiTaiSan"`                                     // Chỉ định loại tài sản
	HanSuDung             *time.Time     `gorm:"column:han_su_dung" json:"hanSuDung"`                                                        // Hạn sử dụng
	Status                *ProductStatus `gorm:"column:status;type:enum('selling','out_of_stock','deleted');default:'selling" json:"status"` // Trạng thái sản phẩm
}

// TableName định nghĩa tên bảng cho model ProductUpdate
func (ProductUpdate) TableName() string { return SanPham{}.TableName() }

// BeforeUpdate là hook được gọi trước khi cập nhật sản phẩm
// Chuyển ProductStatus thành chuỗi để lưu vào database
func (p *ProductUpdate) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.Status != nil {
		tx.Statement.SetColumn("status", p.Status.String())
	}
	return nil
}

// ======================================= INVENTORY =======================================

// KhoHang là struct đại diện cho bảng "kho_hangs" trong database
type KhoHang struct {
	common.SQLModel        // Struct nhúng từ common, chứa các trường ID, CreatedAt, UpdatedAt, DeletedAt
	TenKho          string `gorm:"column:ten_kho" json:"tenKho"` // Tên kho
}

// TonKho là struct đại diện cho bảng "ton_khos" trong database
type TonKho struct {
	common.SQLModel            // Struct nhúng từ common, chứa các trường ID, CreatedAt, UpdatedAt, DeletedAt
	SanPhamID       uint       `gorm:"column:san_pham_id" json:"sanPhamID"`     // ID của sản phẩm
	KhoHangID       uint       `gorm:"column:kho_hang_id" json:"khoHangID"`     // ID của kho hàng
	SoLuong         int        `gorm:"column:so_luong" json:"soLuong"`          // Số lượng tồn kho
	NgayCapNhat     *time.Time `gorm:"column:ngay_cap_nhat" json:"ngayCapNhat"` // Ngày cập nhật tồn kho
	SanPham         SanPham    `gorm:"foreignKey:SanPhamID" json:"sanPham"`     // Sản phẩm (foreign key)
	KhoHang         KhoHang    `gorm:"foreignKey:KhoHangID" json:"khoHang"`     // Kho hàng (foreign key)
}

// DuKienTonKho là struct đại diện cho bảng "du_kien_ton_khos" trong database
type DuKienTonKho struct {
	common.SQLModel            // Struct nhúng từ common, chứa các trường ID, CreatedAt, UpdatedAt, DeletedAt
	MaDuKien        string     `gorm:"column:ma_du_kien;type:varchar(255)" json:"maDuKien"` // Mã dự kiến tồn kho
	SanPhamID       uint       `gorm:"column:san_pham_id" json:"sanPhamID"`                 // ID của sản phẩm
	KhoHangID       uint       `gorm:"column:kho_hang_id" json:"khoHangID"`                 // ID của kho hàng
	SoLuongDuKien   int        `gorm:"column:so_luong_du_kien" json:"soLuongDuKien"`        // Số lượng dự kiến tồn kho
	NgayDuKien      *time.Time `gorm:"column:ngay_du_kien" json:"ngayDuKien"`               // Ngày dự kiến tồn kho
	SanPham         SanPham    `gorm:"foreignKey:SanPhamID" json:"sanPham"`                 // Sản phẩm (foreign key)
	KhoHang         KhoHang    `gorm:"foreignKey:KhoHangID" json:"khoHang"`                 // Kho hàng (foreign key)
}
