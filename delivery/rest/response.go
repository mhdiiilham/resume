package rest

type HTTPResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    any     `json:"data"`
	Error   *string `json:"error"`
}
