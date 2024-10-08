package models

import (
	"time"

	"gorm.io/gorm"
)

type NhomHang struct {
	gorm.Model
	TenNhom string `gorm:"type:varchar(255);not null"`
}

type MucTuHang struct {
	gorm.Model
	TenMucTu string `gorm:"type:varchar(255);not null"`
}

type SanPham struct {
	gorm.Model
	MaHang                        string `gorm:"type:varchar(255);not null"`
	TenSanPham                    string `gorm:"type:varchar(255);not null"`
	NhomHangID                    uint
	MucTuHangID                   uint
	DonViTinh                     string `gorm:"type:varchar(255);not null"`
	VoHieuHoa                     bool
	ChoPhepKhoanThayThe           bool
	DuyTriTonKho                  bool
	BaoQuanCacMatHangTrongSanXuat bool
	CoPhieuMoDau                  bool
	DinhGia                       float64
	TyGiaBanHangTauChuan          float64
	LaChiDinhTaiSan               bool
	HanSuDung                     *time.Time
	NhomHang                      NhomHang  `gorm:"foreignKey:NhomHangID"`
	MucTuHang                     MucTuHang `gorm:"foreignKey:MucTuHangID"`
}

type KhoHang struct {
	gorm.Model
	TenKho string `gorm:"type:varchar(255);not null"`
}

type TonKho struct {
	gorm.Model
	SanPhamID   uint
	KhoHangID   uint
	SoLuong     int
	NgayCapNhat time.Time
	SanPham     SanPham `gorm:"foreignKey:SanPhamID"`
	KhoHang     KhoHang `gorm:"foreignKey:KhoHangID"`
}

type DuKienTonKho struct {
	gorm.Model
	MaDuKien      string `gorm:"type:varchar(255);not null"`
	SanPhamID     uint
	KhoHangID     uint
	SoLuongDuKien int
	NgayDuKien    time.Time
	SanPham       SanPham `gorm:"foreignKey:SanPhamID"`
	KhoHang       KhoHang `gorm:"foreignKey:KhoHangID"`
}

// Struct để trả về dữ liệu sản phẩm và tồn kho
type ProductResponse struct {
	SanPham SanPham  `json:"sanPham"`
	TonKho  []TonKho `json:"tonKho"`
}
