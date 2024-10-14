package model

// Filterr là struct chứa các trường dùng để lọc dữ liệu.
type Filterr struct {
	Status string `json:"status" form:""` // Trường Status dùng để lọc theo trạng thái, nhận giá trị từ JSON ("status") và form data (không có tên trường cụ thể).
}
