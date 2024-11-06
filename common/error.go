package common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppError là struct chứa thông tin lỗi.
type AppError struct {
	StatusCode int         `json:"status_code"`
	RootErr    error       `json:"-"`
	Message    string      `json:"message"`
	Log        string      `json:"log"`
	Key        string      `json:"error_key"`
	Data       interface{} `json:"data"`
	RootError  any
}

func (e *AppError) Error() string {
	return e.Message
}

func NewErrorResponse(root error, msg, log, key string, statusCode int) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot list %s: %v", entity, err),
		fmt.Sprintf("ErrCannotList%s", entity),
		http.StatusBadRequest,
	)
}

// ErrInvalidRequest lỗi request không hợp lệ.
func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request: %v", err.Error(), "ErrInvalidRequest", http.StatusBadRequest)
}

// ErrInternal lỗi server nội bộ.
func ErrInternal(err error) *AppError {
	return NewErrorResponse(err, "Something went wrong in the server", err.Error(), "ErrInternal", http.StatusInternalServerError)
}

// ErrCannotGetEntity lỗi khi không thể lấy entity.
func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("không thể lấy %s", entity),
		fmt.Sprintf("err.cannotGet%s", entity),
		http.StatusBadRequest,
	)
}

// Các hàm lỗi khác (từ code cũ của bạn):
func NewCustomError(root error, msg string, key string, statusCode int) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key, statusCode)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key, statusCode)
}

func NewFullErrorResponse(err error) gin.H { // Giữ nguyên hàm này nếu bạn vẫn cần sử dụng
	res := gin.H{"error": err.Error()}

	if e, ok := err.(*AppError); ok {
		res["error"] = e.Message
		res["root_cause"] = e.RootError
	}

	return res
}

// ErrCannotListEntity, ErrCannotDeleteEntity,... (các hàm lỗi khác - giữ nguyên)

func (e *AppError) GinH() gin.H {
	return gin.H{
		"message": e.Message,
		"error":   e.Key,
		"data":    e.Data,
	}
}
