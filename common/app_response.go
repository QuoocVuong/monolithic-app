package common

// successRes là struct đại diện cho response thành công của API.
type successRes struct {
	Data   interface{} `json:"data"`             // Dữ liệu trả về
	Paging interface{} `json:"paging,omitempty"` // Thông tin phân trang (có thể không có)
	Filter interface{} `json:"filter,omitempty"` // Thông tin lọc (có thể không có)
}

// NewSuccessRespone tạo mới một successRes với dữ liệu, phân trang, và lọc.
func NewSuccessRespone(data, paging, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}

// SimpleSuccessRespone tạo mới một successRes chỉ với dữ liệu, không có phân trang và lọc.
func SimpleSuccessRespone(data interface{}) *successRes {
	return NewSuccessRespone(data, nil, nil)
}
