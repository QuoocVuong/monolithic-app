package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"monolithic-app/models"   // Thay thế bằng đường dẫn thực tế
	"monolithic-app/services" // Thay thế bằng đường dẫn thực tế
)

type DuKienTonKhoController struct {
	DB                  *gorm.DB
	duKienTonKhoService *services.DuKienTonKhoService // Lưu ý sử dụng con trỏ
}

func NewDuKienTonKhoController(db *gorm.DB) *DuKienTonKhoController {
	return &DuKienTonKhoController{
		DB:                  db,
		duKienTonKhoService: services.NewDuKienTonKhoService(db),
	}
}

// API lấy danh sách DuKienTonKho
func (dc *DuKienTonKhoController) GetAllDuKienTonKho(c *gin.Context) {
	duKienTonKho, err := dc.duKienTonKhoService.GetAllDuKienTonKho()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, duKienTonKho)
}

// API lấy DuKienTonKho theo ID
func (dc *DuKienTonKhoController) GetDuKienTonKhoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	duKienTonKho, err := dc.duKienTonKhoService.GetDuKienTonKhoByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "DuKienTonKho not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, duKienTonKho)
}

// API tạo DuKienTonKho mới
func (dc *DuKienTonKhoController) CreateDuKienTonKho(c *gin.Context) {
	var newDuKienTonKho models.DuKienTonKho
	if err := c.ShouldBindJSON(&newDuKienTonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdDuKienTonKho, err := dc.duKienTonKhoService.CreateDuKienTonKho(&newDuKienTonKho)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdDuKienTonKho)
}

// API cập nhật DuKienTonKho
func (dc *DuKienTonKhoController) UpdateDuKienTonKho(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedDuKienTonKho models.DuKienTonKho
	if err := c.ShouldBindJSON(&updatedDuKienTonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedDuKienTonKho.ID = uint(id)

	if err := dc.duKienTonKhoService.UpdateDuKienTonKho(&updatedDuKienTonKho); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDuKienTonKho)
}

// API xóa DuKienTonKho
func (dc *DuKienTonKhoController) DeleteDuKienTonKho(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := dc.duKienTonKhoService.DeleteDuKienTonKho(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DuKienTonKho deleted successfully"})
}
