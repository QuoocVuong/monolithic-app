package model

import (
	"monolithic-app/common"
	"monolithic-app/modules/product/model"
	"time"
)

// Model cho bảng DuKienTonKho
type DuKienTonKho struct {
	common.SQLModel
	MaDuKien      string        `gorm:"column:ma_du_kien;type:varchar(255)" json:"maDuKien"`
	SanPhamID     uint          `gorm:"column:san_pham_id" json:"sanPhamID"`
	KhoHangID     uint          `gorm:"column:kho_hang_id" json:"khoHangID"`
	SoLuongDuKien int           `gorm:"column:so_luong_du_kien" json:"soLuongDuKien"`
	NgayDuKien    *time.Time    `gorm:"column:ngay_du_kien" json:"ngayDuKien"`
	SanPham       model.SanPham `gorm:"foreignKey:SanPhamID" json:"sanPham"`
	KhoHang       KhoHang       `gorm:"foreignKey:KhoHangID" json:"khoHang"`
}

func (DuKienTonKho) TableName() string { return "du_kien_ton_khos" }
