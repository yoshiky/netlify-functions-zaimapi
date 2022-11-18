package codes

type Code string

const (
	OK                  Code = "OK"
	InternalServerError Code = "internal_server_error"
	InvalidParams       Code = "invalid_params"
)
