package model

import (
	"monolithic-app/common"
	"monolithic-app/modules/product/model"
	"time"
)

// Model cho bảng TonKho
type TonKho struct {
	common.SQLModel
	SanPhamID   uint          `gorm:"column:san_pham_id" json:"sanPhamID"`
	KhoHangID   uint          `gorm:"column:kho_hang_id" json:"khoHangID"`
	SoLuong     int           `gorm:"column:so_luong" json:"soLuong"`
	NgayCapNhat *time.Time    `gorm:"column:ngay_cap_nhat" json:"ngayCapNhat"`
	SanPham     model.SanPham `gorm:"foreignKey:SanPhamID" json:"sanPham"`
	KhoHang     KhoHang       `gorm:"foreignKey:KhoHangID" json:"khoHang"`
}

func (TonKho) TableName() string { return "ton_khos" }
