package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type JSONRespBody struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

// StatusOK for 200
func StatusOK(w http.ResponseWriter, data interface{}) {
	body := &JSONRespBody{
		StatusCode: http.StatusOK,
		Data:       data,
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// StatusCreated for 201
func StatusCreated(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusCreated, // Set the status to 201 Created
		Data:       "created",
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated) // Set HTTP status to 201
	w.Write(b)
}

// StatusNoContent for 204
func StatusNoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent) // Set HTTP status to 201
}

// BadRequest for 400
func BadRequest(w http.ResponseWriter, err error) {
	body := &JSONRespBody{
		StatusCode: http.StatusBadRequest,
		Error:      err.Error(),
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(b)
}

// NotFound for 404
func NotFound(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusNotFound,
		Error:      "resource not found",
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write(b)
}

// StatusConflict for 409
func StatusConflict(w http.ResponseWriter, msg string) {
	body := &JSONRespBody{
		StatusCode: http.StatusConflict,
		Data:       msg,
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusConflict)
	w.Write(b)
}

// UnprocessableEntity for 422
func UnprocessableEntity(w http.ResponseWriter, errors error) {
	var err interface{}
	if fieldErrors, ok := errors.(validator.ValidationErrors); ok {
		fieldErrorsMap := make(map[string]string, len(fieldErrors))
		for _, ve := range fieldErrors {
			fieldErrorsMap[ve.Namespace()] = ve.Translate(nil)
		}
		err = fieldErrorsMap
	} else {
		err = errors.Error()
	}

	body := &JSONRespBody{
		StatusCode: http.StatusUnprocessableEntity,
		Error:      err,
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(b)
}

// InternalServerError for 500
func InternalServerError(w http.ResponseWriter) {
	body := &JSONRespBody{
		StatusCode: http.StatusInternalServerError,
		Data:       "Internal Server Error",
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error during marshalling json body")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}
