package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ProductStatus int

const (
	ProductStatusSelling ProductStatus = iota
	ProductStatusOutOfStock
	ProductStatusDeleted
)

var allProductStatus = [3]string{"selling", "out_of_stock", "deleted"}

func (product ProductStatus) String() string {
	return allProductStatus[int(product)]

}
func parseStrtoProducStatus(s string) (ProductStatus, error) {
	for i := range allProductStatus {
		if allProductStatus[i] == s {
			return ProductStatus(i), nil
		}
	}
	return ProductStatus(0), errors.New("invalid product status")
}

//	func (product *ProductStatus) scan(value interface{}) error {
//		bytes, ok := value.([]byte)
//		if !ok {
//			return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
//		}
//
//		v, err := parseStrtoProducStatus(string(bytes))
//		if err != nil {
//			return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
//		}
//		*product = v
//		return nil
//	}
func (product *ProductStatus) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	// Chuyển đổi []byte sang string trước khi parse
	byteVal, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan ProductStatus: %v", value)
	}
	strVal := string(byteVal)

	v, err := parseStrtoProducStatus(strVal)
	if err != nil {
		return fmt.Errorf("failed to parse ProductStatus: %v", err)
	}

	*product = v
	return nil
}

func (product *ProductStatus) Value() (driver.Value, error) {
	if product == nil {
		return nil, nil
	}

	return product.String(), nil
}
func (product *ProductStatus) MarshalJSON() ([]byte, error) {
	if product == nil {
		return nil, nil
	}
	//return []byte(fmt.Sprintf("\"%s\"", product.String())), nil
	return json.Marshal(product.String())
}

//	func (product *ProductStatus) UnmarshalJSON(data []byte) error {
//		str := strings.ReplaceAll(string(data), "\"", "")
//		productValue, err := parseStrtoProducStatus(str)
//		if err != nil {
//			return err
//		}
//		*product = productValue
//		return nil
//	}
func (product *ProductStatus) UnmarshalJSON(data []byte) error {
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return fmt.Errorf("failed to unmarshal ProductStatus: %w", err)
	}

	v, err := parseStrtoProducStatus(strVal)
	if err != nil {
		return fmt.Errorf("failed to parse ProductStatus: %w", err)
	}

	*product = v
	return nil
}
