package model

import (
	"errors"
	"monolithic-app/common"
	"time"
)

var (
	ErrMaHangIsBlank  = errors.New("MaHang CanNot Is Blank")
	ErrProductDeleted = errors.New("product is deleted")
)

type NhomHang struct {
	common.SQLModel
	TenNhom string `gorm:"column:ten_nhom;type:varchar(255)" json:"tenNhom"`
}

type SanPham struct {
	common.SQLModel
	MaHang                string         `gorm:"column:ma_hang;type:varchar(255)" json:"maHang"`
	NhomHangID            uint           `gorm:"column:nhom_hang_id;index;foreignKey:NhomHangID" json:"nhomHangID"`
	TenSanPham            string         `gorm:"column:ten_san_pham;type:varchar(255)" json:"tenSanPham"`
	DonViTinh             string         `gorm:"column:don_vi_tinh;type:varchar(255)" json:"donViTinh"`
	CoPhieuMoDau          float64        `gorm:"column:co_phieu_mo_dau" json:"coPhieuMoDau"`
	DinhGia               float64        `gorm:"column:dinh_gia" json:"dinhGia"`
	TyGiaBanHangTieuChuan float64        `gorm:"column:ty_gia_ban_hang_tieu_chuan" json:"tyGiaBanHangTieuChuan"`
	ChiDinhLoaiTaiSan     string         `gorm:"column:chi_dinh_loai_tai_san;type:varchar(255)" json:"ChiDinhLoaiTaiSan"`
	Status                *ProductStatus `gorm:"column:status;type:enum('selling','out_of_stock','deleted');default:'selling" json:"status"`
	HanSuDung             *time.Time     `gorm:"column:han_su_dung" json:"hanSuDung"`
	NhomHang              NhomHang       `gorm:"foreignKey:NhomHangID" json:"nhomHang"`
}

func (SanPham) TableName() string { return "san_phams" }

type ProductCreation struct {
	Id                    int            `gorm:"column:id" json:"-"`
	MaHang                string         `gorm:"column:ma_hang;type:varchar(255)" json:"maHang"`
	TenSanPham            string         `gorm:"column:ten_san_pham;type:varchar(255)" json:"tenSanPham"`
	DonViTinh             string         `gorm:"column:don_vi_tinh;type:varchar(255)" json:"donViTinh"`
	CoPhieuMoDau          float64        `gorm:"column:co_phieu_mo_dau" json:"coPhieuMoDau"`
	DinhGia               float64        `gorm:"column:dinh_gia" json:"dinhGia"`
	TyGiaBanHangTieuChuan float64        `gorm:"column:ty_gia_ban_hang_tieu_chuan" json:"tyGiaBanHangTieuChuan"`
	ChiDinhLoaiTaiSan     string         `gorm:"column:chi_dinh_loai_tai_san;type:varchar(255)" json:"ChiDinhLoaiTaiSan"`
	Status                *ProductStatus `gorm:"column:status;type:enum('selling','out_of_stock','deleted');default:'selling" json:"status"`
}

func (ProductCreation) TableName() string { return SanPham{}.TableName() }

// dung con tro de update duoc chuoi rong
type ProductUpdate struct {
	MaHang                *string        `gorm:"column:ma_hang " json:"maHang"`
	TenSanPham            *string        `gorm:"column:ten_san_pham " json:"tenSanPham"`
	DonViTinh             *string        `gorm:"column:don_vi_tinh " json:"donViTinh"`
	CoPhieuMoDau          *string        `gorm:"column:co_phieu_mo_dau " json:"coPhieuMoDau"`
	DinhGia               *string        `gorm:"column:dinh_gia " json:"dinhGia"`
	TyGiaBanHangTieuChuan *float64       `gorm:"column:ty_gia_ban_hang_tieu_chuan" json:"tyGiaBanHangTieuChuan"`
	ChiDinhLoaiTaiSan     *string        `gorm:"column:chi_dinh_loai_tai_san " json:"ChiDinhLoaiTaiSan"`
	Status                *ProductStatus `gorm:"column:status;type:enum('selling','out_of_stock','deleted');default:'selling" json:"status"`
}

func (ProductUpdate) TableName() string { return SanPham{}.TableName() }

type KhoHang struct {
	common.SQLModel
	TenKho string `gorm:"column:ten_kho" json:"tenKho"`
}
type TonKho struct {
	common.SQLModel
	SanPhamID   uint       `gorm:"column:san_pham_id" json:"sanPhamID"`
	KhoHangID   uint       `gorm:"column:kho_hang_id" json:"khoHangID"`
	SoLuong     int        `gorm:"column:so_luong" json:"soLuong"`
	NgayCapNhat *time.Time `gorm:"column:ngay_cap_nhat" json:"ngayCapNhat"`
	SanPham     SanPham    `gorm:"foreignKey:SanPhamID" json:"sanPham"`
	KhoHang     KhoHang    `gorm:"foreignKey:KhoHangID" json:"khoHang"`
}
type DuKienTonKho struct {
	common.SQLModel
	MaDuKien      string     `gorm:"column:ma_du_kien;type:varchar(255)" json:"maDuKien"`
	SanPhamID     uint       `gorm:"column:san_pham_id" json:"sanPhamID"`
	KhoHangID     uint       `gorm:"column:kho_hang_id" json:"khoHangID"`
	SoLuongDuKien int        `gorm:"column:so_luong_du_kien" json:"soLuongDuKien"`
	NgayDuKien    *time.Time `gorm:"column:ngay_du_kien" json:"ngayDuKien"`
	SanPham       SanPham    `gorm:"foreignKey:SanPhamID" json:"sanPham"`
	KhoHang       KhoHang    `gorm:"foreignKey:KhoHangID" json:"khoHang"`
}

// Struct để trả về dữ liệu sản phẩm và tồn kho
type ProductResponse struct {
	SanPham SanPham  `json:"sanPham"`
	TonKho  []TonKho `json:"tonKho"`
}

//func (product *ProductStatus) MarshalJSON() ([]byte, error) {
//	return []byte(fmt.Sprintf("\"%s\""), product.String())), nil
//}
