package response

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	Err() (err error)
	JSON(w http.ResponseWriter) (err error)
}

type responseImpl struct {
	err    error
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(status string, data interface{}) (resp Response) {
	return &responseImpl{
		err:    nil,
		Status: status,
		Data:   data,
	}
}

func Error(status string, err error) (resp Response) {
	return &responseImpl{
		err:    err,
		Status: status,
		Data:   nil,
	}
}

func (r *responseImpl) getStatusCode(status string) (statusCode int) {
	switch status {
	case StatusOK:
		return http.StatusOK
	case StatusCreated:
		return http.StatusCreated
	case StatusBadRequest:
		return http.StatusBadRequest
	case StatusUnauthorized:
		return http.StatusUnauthorized
	case StatusForbiddend:
		return http.StatusForbidden
	case StatusNotFound:
		return http.StatusNotFound
	case StatusConflicted:
		return http.StatusConflict
	case StatusUnprocessableEntity:
		return http.StatusUnprocessableEntity
	case StatusInternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func (r *responseImpl) Err() (err error) {
	return r.err
}

func (r *responseImpl) JSON(w http.ResponseWriter) error {
	statusCode := r.getStatusCode(r.Status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(r)
}
