package types

type Success struct {
	Message string `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}
