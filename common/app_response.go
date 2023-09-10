package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
