package response

type (
	Response struct {
		Code int `json:"-"`

		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	ResponseError interface {
		Responses()
		Error() string
	}
)
