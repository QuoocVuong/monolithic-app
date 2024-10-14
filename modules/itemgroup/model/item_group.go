package model

import (
	"errors"
	"monolithic-app/common"
)

type NhomHang struct {
	common.SQLModel
	TenNhom string `gorm:"column:ten_nhom;type:varchar(255)" json:"tenNhom"`
}

func (NhomHang) TableName() string { return "nhom_hangs" }

type ItemGroupCreation struct {
	common.SQLModel
	TenNhom string `gorm:"column:ten_nhom;type:varchar(255)" json:"tenNhom"`
}

func (ItemGroupCreation) TableName() string { return NhomHang{}.TableName() }

type ItemGroupUpdate struct {
	common.SQLModel
	TenNhom *string `gorm:"column:ten_nhom;type:varchar(255)" json:"tenNhom"`
}

func (ItemGroupUpdate) TableName() string { return NhomHang{}.TableName() }

var (
	ErrItemGroupIsBlank = errors.New("Nhom Hang CanNot Is Blank")
	ErrItemGroupDeleted = errors.New("Nhom Hang is deleted")
)
