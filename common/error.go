package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// BaseError là struct chứa thông tin lỗi cơ bản.
type BaseError struct {
	StatusCode int    `json:"status_code"`
	RootError  error  `json:"-"` // Không serialize RootError
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

// Error trả về thông báo lỗi.
func (e BaseError) Error() string {
	return e.Message
}

// NewBaseError tạo một BaseError mới.
func NewBaseError(root error, msg string, log string, key string) *BaseError {
	return &BaseError{
		RootError: root,
		Message:   msg,
		Log:       log,
		Key:       key,
	}
}

// NewCustomError tạo một lỗi custom với message và key.
func NewCustomError(root error, message string, key string) *BaseError {
	if root != nil {
		return NewBaseError(root, message, root.Error(), key)
	}

	return NewBaseError(errors.New(message), message, message, key)
}

// NewFullErrorResponse tạo một response lỗi đầy đủ.
func NewFullErrorResponse(err error) gin.H {
	res := gin.H{"error": err.Error()}

	if e, ok := err.(*BaseError); ok {
		res["error"] = e.Message
		res["root_cause"] = e.RootError
	}

	return res
}

// ErrCannotGetEntity lỗi khi không thể lấy entity.
func ErrCannotGetEntity(entity string, err error) *BaseError {
	return NewCustomError(
		err,
		fmt.Sprintf("không thể lấy %s", entity),
		fmt.Sprintf("err.cannotGet%s", entity),
	)
}

// ... (các hàm error khác trong common package)
