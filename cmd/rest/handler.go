package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"capital-gains-api/internal/entity"
	"capital-gains-api/internal/service/operation"
)

type Handler struct {
	controller       IController
	operationService operation.IService
}

type Response struct {
	Content any    `json:"content"`
	Message string `json:"message"`
	Code    int    `json:code`
}

func NewHandler(controller IController, operationService operation.IService) *Handler {
	return &Handler{
		controller:       controller,
		operationService: operationService,
	}
}

func (h *Handler) PostTaxOperation(w http.ResponseWriter, r *http.Request) {

	var operations []entity.Operation

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(operations); err != nil {
		fmt.Println(err, "error converting tax list to json")
		return
	}

	response, err := h.controller.PostTaxOperation(r.Context(), operations)
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, err.Error(), nil)
	}

	h.writeResponse(w, http.StatusInternalServerError, "", response)

}

func (h *Handler) writeResponse(w http.ResponseWriter, statusCode int, msg string, content any) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := &Response{
		Code:    statusCode,
		Message: msg,
		Content: content,
	}

	json.NewEncoder(w).Encode(response)

}
