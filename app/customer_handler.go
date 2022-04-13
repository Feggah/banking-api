package app

import (
	"encoding/json"
	"net/http"

	"github.com/Feggah/banking-api/errors"
	"github.com/Feggah/banking-api/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (c *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := c.service.GetAllCustomer()
	if err != nil {
		appErr, _ := err.(*errors.AppErr)
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customers)
}

func (c *CustomerHandlers) getCustomersByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.FormValue("status")
	if status != "active" && status != "inactive" {
		appErr := errors.NewBadRequestError("status can only be filtered by \"active\" or \"inactive\"")
		writeResponse(w, http.StatusBadRequest, appErr.AsMessage())
		return
	}

	customers, err := c.service.GetAllCustomersByStatus(status)
	if err != nil {
		appErr, _ := err.(*errors.AppErr)
		writeResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (c *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := c.service.GetCustomer(id)
	if err != nil {
		appErr, _ := err.(*errors.AppErr)
		writeResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
