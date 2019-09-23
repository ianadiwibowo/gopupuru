package handler

import (
	"encoding/json"
	"net/http"
)

// Response is the standard API data-meta response
type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

// Message is used only for errors
type Message struct {
	Message string `json:"message"`
}

// Meta is used to explain Data
type Meta struct {
	HTTPStatus int `json:"http_status"`
}

// Success 200
func respondWithSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusOK
	w.WriteHeader(httpStatus)

	response := Response{
		Data: data,
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// Not Found 404
func respondWithNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusNotFound
	w.WriteHeader(httpStatus)

	response := Response{
		Data: Message{
			Message: "Not found",
		},
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// Unprocessable Entity 422
func respondWithUnprocessableEntity(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusUnprocessableEntity
	w.WriteHeader(httpStatus)

	response := Response{
		Data: Message{
			Message: "Unprocessable entity",
		},
		Meta: Meta{
			HTTPStatus: httpStatus,
		},
	}

	json.NewEncoder(w).Encode(response)
}
