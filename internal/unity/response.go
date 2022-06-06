package unity

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
