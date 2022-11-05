package commons

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//error response struct
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}