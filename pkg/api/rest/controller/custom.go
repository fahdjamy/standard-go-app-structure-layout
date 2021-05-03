package controller

import (
	"encoding/json"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	logger "github.com/rs/zerolog/log"
	"log"
	"net/http"
)

type CustomController struct {
	logger *log.Logger
}

func (c *CustomController) PageNotFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	path := request.RequestURI

	response := types.Success{
		Data: nil,
		Success: false,
		Message: "Page '" + path + "' not found",
	}

	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		logger.Print(err)
	}
}

func NewCustomController( logger *log.Logger) CustomController {
	return CustomController{ logger}
}
