package model

import (
	"errors"
	"monolithic-app/common"
)

// Model cho bảng KhoHang
type KhoHang struct {
	common.SQLModel
	TenKho string `gorm:"column:ten_kho" json:"tenKho"`
}

func (KhoHang) TableName() string { return "kho_hangs" }

// Các struct cho Create, Update, Filter, ... (nếu cần)
// Ví dụ:
type KhoHangCreate struct {
	common.SQLModel
	TenKho string `json:"tenKho" binding:"required"` // required: bắt buộc phải có khi tạo
}

type KhoHangUpdate struct {
	common.SQLModel
	TenKho *string `json:"tenKho" ` //  ko bắt buộc phải có khi update
}

var (
	ErrKhoHangExisted = errors.New("tên kho hàng đã tồn tại")
)
