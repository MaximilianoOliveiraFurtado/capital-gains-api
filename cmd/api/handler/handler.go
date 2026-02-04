package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"capital-gains-api/internal/entity"
	"capital-gains-api/internal/service/operation"

	"capital-gains-api/cmd/api/controller"
)

type Handler struct {
	controller       controller.IController
	operationService operation.IService
}

type Response struct {
	Content any    `json:"content"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewHandler(controller controller.IController, operationService operation.IService) *Handler {
	return &Handler{
		controller:       controller,
		operationService: operationService,
	}
}

func (h *Handler) PostTaxOperation(w http.ResponseWriter, r *http.Request) {

	var operationsInput []entity.Operation

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&operationsInput); err != nil {
		fmt.Println(err, "error converting tax list to json")
		return
	}

	response, err := h.controller.PostTaxOperation(r.Context(), operationsInput)
	if err != nil {
		h.writeResponse(w, http.StatusInternalServerError, err.Error(), nil)
	}

	h.writeResponse(w, http.StatusOK, "Ok", response)

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
