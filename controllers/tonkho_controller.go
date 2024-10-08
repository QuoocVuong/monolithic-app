package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"monolithic-app/models"   // Thay thế bằng đường dẫn thực tế
	"monolithic-app/services" // Thay thế bằng đường dẫn thực tế
)

type TonKhoController struct {
	DB            *gorm.DB
	tonKhoService *services.TonKhoService // Sử dụng con trỏ
}

func NewTonKhoController(db *gorm.DB) *TonKhoController {
	return &TonKhoController{
		DB:            db,
		tonKhoService: services.NewTonKhoService(db),
	}
}

// API lấy danh sách TonKho
func (tc *TonKhoController) GetAllTonKho(c *gin.Context) {
	tonKho, err := tc.tonKhoService.GetAllTonKho()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tonKho)
}

// API lấy TonKho theo ID
func (tc *TonKhoController) GetTonKhoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	tonKho, err := tc.tonKhoService.GetTonKhoByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "TonKho not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, tonKho)
}

// API tạo TonKho mới
func (tc *TonKhoController) CreateTonKho(c *gin.Context) {
	var newTonKho models.TonKho
	if err := c.ShouldBindJSON(&newTonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTonKho, err := tc.tonKhoService.CreateTonKho(&newTonKho)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTonKho)
}

// API cập nhật TonKho
func (tc *TonKhoController) UpdateTonKho(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTonKho models.TonKho
	if err := c.ShouldBindJSON(&updatedTonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTonKho.ID = uint(id)

	if err := tc.tonKhoService.UpdateTonKho(&updatedTonKho); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTonKho)
}

// API xóa TonKho
func (tc *TonKhoController) DeleteTonKho(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := tc.tonKhoService.DeleteTonKho(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TonKho deleted successfully"})
}
