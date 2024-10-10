package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
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
func (product *ProductStatus) scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	v, err := parseStrtoProducStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
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
	return []byte(fmt.Sprintf("\"%s\"", product.String())), nil
}
func (product *ProductStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	productValue, err := parseStrtoProducStatus(str)
	if err != nil {
		return err
	}
	*product = productValue
	return nil
}
