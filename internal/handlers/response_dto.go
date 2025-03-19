package handlers

type BaseRespDTO struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// type FailedRespDTO struct {
// 	BaseRespDTO
// 	Error string `json:"error"`
// }

// type SuccessRespDTO struct {
// 	BaseRespDTO
// }
