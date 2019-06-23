package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/diegogomesaraujo/central-loteria/pkg/exception"
)

func readBodyFromJSON(w http.ResponseWriter, r *http.Request, entity interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		handleExceptionError(w, "Content-Type not supported", http.StatusUnsupportedMediaType)
		return fmt.Errorf("Content-Type not supported")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error when read request body: %v\n", err)
		handleExceptionError(w, "Erro when try read body", http.StatusInternalServerError)
		return fmt.Errorf("Erro when try read body")
	}

	if !json.Valid(body) {
		handleExceptionError(w, "Bad Request", http.StatusBadRequest)
		return fmt.Errorf("Bad Request")
	}

	json.Unmarshal(body, entity)

	return nil
}

func handleExceptionError(w http.ResponseWriter, message string, httpCode int) {
	ex := exception.Exception{
		Message: message,
		Code:    httpCode,
	}
	exception.HandleError(w, ex)
}
