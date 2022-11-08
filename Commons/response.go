package commons

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}

// error response struct
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
