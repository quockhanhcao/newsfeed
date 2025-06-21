package http

type ResponseCode int

const (
	Success ResponseCode = iota
	BadRequest
	Unauthorized
	Forbidden
	InternalServerError
)

type CommonResponse struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data,omitempty"`
}
