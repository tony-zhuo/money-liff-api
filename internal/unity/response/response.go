package response

import "net/http"

// OkResponse 200, 201...
type OkResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Exception struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error is required by the error interface.
func (e Exception) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e Exception) StatusCode() int {
	return e.Code
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

// BadRequest 400 error
func BadRequest(msg string) Exception {
	if msg == "" {
		msg = "Request field format not correct."
	}
	return Exception{
		Status:  false,
		Code:    400,
		Message: msg,
	}
}

// Unauthorized 401 error
func Unauthorized(msg string) Exception {
	if msg == "" {
		msg = "Unauthorized."
	}
	return Exception{
		Status:  false,
		Code:    401,
		Message: msg,
	}
}

// NotFound 404 error
func NotFound(msg string) Exception {
	if msg == "" {
		msg = "Not found."
	}
	return Exception{
		Status:  false,
		Code:    404,
		Message: msg,
	}
}

// InternalServerError 500 error
func InternalServerError(msg string) Exception {
	if msg == "" {
		msg = "Server error."
	}
	return Exception{
		Status:  false,
		Code:    500,
		Message: msg,
	}
}
