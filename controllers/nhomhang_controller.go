package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"monolithic-app/models"   // Thay thế bằng đường dẫn thực tế
	"monolithic-app/services" // Thay thế bằng đường dẫn thực tế
)

type NhomHangController struct {
	DB              *gorm.DB
	nhomHangService *services.NhomHangService // Sử dụng con trỏ
}

func NewNhomHangController(db *gorm.DB) *NhomHangController {
	return &NhomHangController{
		DB:              db,
		nhomHangService: services.NewNhomHangService(db),
	}
}

// API lấy danh sách NhomHang
func (nc *NhomHangController) GetAllNhomHang(c *gin.Context) {
	nhomHang, err := nc.nhomHangService.GetAllNhomHang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nhomHang)
}

// API lấy NhomHang theo ID
func (nc *NhomHangController) GetNhomHangByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	nhomHang, err := nc.nhomHangService.GetNhomHangByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "NhomHang not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, nhomHang)
}

// API tạo NhomHang mới
func (nc *NhomHangController) CreateNhomHang(c *gin.Context) {
	var newNhomHang models.NhomHang
	if err := c.ShouldBindJSON(&newNhomHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdNhomHang, err := nc.nhomHangService.CreateNhomHang(&newNhomHang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdNhomHang)
}

// API cập nhật NhomHang
func (nc *NhomHangController) UpdateNhomHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedNhomHang models.NhomHang
	if err := c.ShouldBindJSON(&updatedNhomHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedNhomHang.ID = uint(id)

	if err := nc.nhomHangService.UpdateNhomHang(&updatedNhomHang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedNhomHang)
}

// API xóa NhomHang
func (nc *NhomHangController) DeleteNhomHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := nc.nhomHangService.DeleteNhomHang(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "NhomHang deleted successfully"})
}
