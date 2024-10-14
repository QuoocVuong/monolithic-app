package common

// Paging là struct chứa thông tin phân trang.
type Paging struct {
	Page  int    `json:"page" form:"page"`   // Trang hiện tại (lấy từ query parameter "page")
	Limit int    `json:"limit" form:"limit"` // Số lượng bản ghi trên mỗi trang (lấy từ query parameter "limit")
	Total int64  `json:"total" form:"-"`     // Tổng số bản ghi (không lấy từ form data)
	Sort  string `json:"sort" form:"sort"`   // Trường sắp xếp (lấy từ query parameter "sort")
}

// Process xử lý và validate thông tin phân trang.
func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1 // Nếu Page <= 0, mặc định là trang 1
	}
	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 5 // Nếu Limit <= 0 hoặc Limit >= 100, mặc định là 5 bản ghi/trang
	}
}
