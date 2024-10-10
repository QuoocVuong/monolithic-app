package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessRespone(data, paging, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}
func SimpleSuccessRespone(data interface{}) *successRes {
	return NewSuccessRespone(data, nil, nil)

}
