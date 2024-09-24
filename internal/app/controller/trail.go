package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sebastianmendez/trail-browser/internal/app/controller/utils"
	"github.com/sebastianmendez/trail-browser/internal/app/service"
)

// this layer will take the request, validate and process it and send it sanitized to the service layer
func HandleList(w http.ResponseWriter, r *http.Request) {

	filters := utils.ParseFilters(r.URL.RawQuery)

	trails, err := service.List(r.Context(), filters)
	if err != nil {
		httpWriteResponse(w, nil, http.StatusUnprocessableEntity)
		return
	}

	body, code := buildResponse(http.StatusOK, trails)
	httpWriteResponse(w, body, code)
}

func httpWriteResponse(w http.ResponseWriter, body []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Printf("error %v", err)
	}
}

func buildResponse(statusCode int, body interface{}) ([]byte, int) {
	bytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("failed to json marshal response: %v with error: %v", body, err)

		return nil, http.StatusInternalServerError
	}

	return bytes, statusCode
}
