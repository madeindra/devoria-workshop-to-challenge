package response

import (
	"encoding/json"
	"net/http"
)

// TODO: Add http status code
type Response interface {
	JSON(w http.ResponseWriter) (err error)
}

type responseImpl struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}) (resp Response) {
	return &responseImpl{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func Error(err error) (resp Response) {
	return &responseImpl{
		Status:  false,
		Message: err.Error(),
		Data:    nil,
	}
}

func (r *responseImpl) JSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(r)
}
