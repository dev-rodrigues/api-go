package dto

type (
	// Response struct for REST API
	Response struct {
		Response Rest        `json:"response"`
		Data     interface{} `json:"data,omitempty"`
	}
	// Rest struct for response Res API
	Rest struct {
		Message string `json:"message,omitempty"`
		Code    string `json:"code,omitempty"`
	}
)
