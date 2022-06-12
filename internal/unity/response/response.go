package response

import "net/http"

// OkResponse 200, 201...
type OkResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponse struct {
	Status  bool       `json:"status"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    Pagination `json:"data"`
}

type Pagination struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	TotalPage  int         `json:"total_page"`
	TotalCount int         `json:"total_count"`
	Result     interface{} `json:"result"`
}

func Ok(msg string, data interface{}) OkResponse {
	if msg == "" {
		msg = "Success."
	}
	return OkResponse{
		Status:  true,
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func Created(msg string) OkResponse {
	if msg == "" {
		msg = "Created success."
	}
	return OkResponse{
		Status:  true,
		Code:    http.StatusCreated,
		Message: msg,
		Data:    nil,
	}
}

// List 200 分頁 response
func List(p Pagination, msg string) ListResponse {
	if msg == "" {
		msg = "Success."
	}
	return ListResponse{
		Status:  true,
		Code:    http.StatusOK,
		Message: msg,
		Data:    p,
	}
}
