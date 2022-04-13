package app

import (
	"encoding/json"
	"net/http"

	"github.com/Feggah/banking-api/dto"
	"github.com/Feggah/banking-api/errors"
	"github.com/Feggah/banking-api/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerID = customerID
	account, err := h.service.NewAccount(request)
	if err != nil {
		appErr, _ := err.(*errors.AppErr)
		writeResponse(w, appErr.Code, appErr.Message)
		return
	}

	writeResponse(w, http.StatusCreated, account)
}
