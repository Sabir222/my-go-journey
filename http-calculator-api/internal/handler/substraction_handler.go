package handler

import (
	"encoding/json"
	"os"
	//"io"
	"github.com/go-playground/validator/v10"
	//"log"
	"fmt"
	"log/slog"
	"net/http"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type InputData struct {
	Numbers []int `json:"Numbers" validate:"required,dive,gt=0"`
}

func SubstractionHandler(w http.ResponseWriter, r *http.Request) {

	logger.Info("Received Request",
		"method", r.Method,
		"path", r.URL.Path,
		"ip", r.RemoteAddr)

	defer r.Body.Close()

	BodyData := InputData{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&BodyData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(&BodyData); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	if len(BodyData.Numbers) == 0 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := BodyData.Numbers[0]

	for _, number := range BodyData.Numbers[1:] {
		result -= number
	}
	response := map[string]int{"Result": result}
	jsonResp, err := json.Marshal(response)
	// Marshal error handling
	if err != nil {
		http.Error(w, "Error sendfing back results", http.StatusInternalServerError)
		return
	}

	//Response Error handling
	if _, err = w.Write(jsonResp); err != nil {
		http.Error(w, "Error sending back results", http.StatusInternalServerError)
		return
	}

	logger.Info("Request processed successfully",
		"result", result,
	)
	w.WriteHeader(http.StatusOK)
}
