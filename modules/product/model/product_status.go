package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// ProductStatus là một kiểu enum đại diện cho trạng thái của sản phẩm.
type ProductStatus int

// Định nghĩa các giá trị cho ProductStatus.
const (
	ProductStatusSelling    ProductStatus = iota // 0: Sản phẩm đang bán
	ProductStatusOutOfStock                      // 1: Sản phẩm hết hàng
	ProductStatusDeleted                         // 2: Sản phẩm đã bị xóa
)

// allProductStatus là một mảng chứa tất cả các giá trị chuỗi của ProductStatus.
var allProductStatus = [3]string{"selling", "out_of_stock", "deleted"}

// String chuyển đổi ProductStatus thành chuỗi.
func (product ProductStatus) String() string {
	return allProductStatus[int(product)]
}

// parseStrtoProducStatus chuyển đổi một chuỗi thành ProductStatus.
func parseStrtoProducStatus(s string) (ProductStatus, error) {
	for i := range allProductStatus {
		if allProductStatus[i] == s {
			return ProductStatus(i), nil
		}
	}
	return ProductStatus(0), errors.New("invalid product status") // Trả về lỗi nếu chuỗi không hợp lệ
}

// Scan implement interface sql.Scanner để đọc dữ liệu từ database.
func (product *ProductStatus) Scan(value interface{}) error {
	if value == nil {
		return nil // Không có lỗi nếu value là nil
	}

	// Chuyển đổi value (interface{}) sang []byte
	byteVal, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan ProductStatus: %v", value) // Trả về lỗi nếu chuyển đổi thất bại
	}

	// Chuyển đổi []byte sang string
	strVal := string(byteVal)

	// Chuyển đổi string sang ProductStatus
	v, err := parseStrtoProducStatus(strVal)
	if err != nil {
		return fmt.Errorf("failed to parse ProductStatus: %v", err) // Trả về lỗi nếu parse thất bại
	}

	*product = v // Gán giá trị cho product
	return nil   // Trả về nil nếu thành công
}

// Value implement interface driver.Valuer để ghi dữ liệu vào database.
func (product *ProductStatus) Value() (driver.Value, error) {
	if product == nil {
		return nil, nil // Trả về nil nếu product là nil
	}

	return product.String(), nil // Trả về chuỗi đại diện cho ProductStatus
}

// MarshalJSON implement interface json.Marshaler để chuyển đổi ProductStatus sang JSON.
func (product *ProductStatus) MarshalJSON() ([]byte, error) {
	if product == nil {
		return nil, nil // Trả về nil nếu product là nil
	}

	return json.Marshal(product.String()) // Chuyển đổi chuỗi đại diện cho ProductStatus sang JSON
}

// UnmarshalJSON implement interface json.Unmarshaler để chuyển đổi JSON sang ProductStatus.
func (product *ProductStatus) UnmarshalJSON(data []byte) error {
	var strVal string // Biến lưu trữ giá trị chuỗi từ JSON

	// Chuyển đổi JSON sang string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return fmt.Errorf("failed to unmarshal ProductStatus: %w", err) // Trả về lỗi nếu unmarshal thất bại
	}

	// Chuyển đổi string sang ProductStatus
	v, err := parseStrtoProducStatus(strVal)
	if err != nil {
		return fmt.Errorf("failed to parse ProductStatus: %w", err) // Trả về lỗi nếu parse thất bại
	}

	*product = v // Gán giá trị cho product
	return nil   // Trả về nil nếu thành công
}
