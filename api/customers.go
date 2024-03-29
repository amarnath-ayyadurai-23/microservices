package api

import (
	"encoding/json"
	"net/http"

	"github.com/amarnath-ayyadurai-23/microservices/models/customers"
	"github.com/dimfeld/httptreemux/v5"
)

// type JsCustomer struct {
// 	ID        string `json:"customer_id"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Email     string `json:"email"`
// 	Phone     string `json:"phone"`
// 	Address   string `json:"address"`
// }

// func toJson(dbRS customers.DBCustomer) JsCustomer {
// 	p := (*JsCustomer)(unsafe.Pointer(&dbRS))
// 	return *p
// }

// func toJsonSlice(dbSRs []customers.DBCustomer) []JsCustomer {
// 	rs := make([]JsCustomer, len(dbSRs))
// 	for i, dbSR := range dbSRs {
// 		rs[i] = toJson(dbSR)
// 	}

// 	return rs
// }
func (a *api) GetCustomers(w http.ResponseWriter, r *http.Request) {
		
		w.Header().Set("Content-Type", "application/json")
		cus, err := a.customer.GetCustomers()
		if err != nil {
			a.log.Printf("<Database> %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(cus)
		if err != nil {
			a.log.Printf("<API> %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
}

func (a *api) GetCustomer(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	id := httptreemux.ContextParams(r.Context())["id"]
	
	cus, err := a.customer.GetCustomer(id)
	if err != nil {
		a.log.Printf("<Database> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(cus)
	if err != nil {
		a.log.Printf("<API> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *api) CreateCustomer(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	var cus customers.DBCustomer
	err := json.NewDecoder(r.Body).Decode(&cus)
	if err != nil {
		a.log.Printf("<API> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = a.customer.CreateCustomer(cus)
	if err != nil {
		a.log.Printf("<Database> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	
}

func (a *api) QuerybyEmail(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	email := httptreemux.ContextParams(r.Context())["email"]
	
	cus, err := a.customer.QuerybyEmail(email)
	if err != nil {
		a.log.Printf("<Database> %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if cus == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(cus)
	if err != nil {
		a.log.Printf("<API> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *api) DeleteCustomer(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	id := httptreemux.ContextParams(r.Context())["id"]
	
	err := a.customer.DeletebyID(id)
	if err != nil {
		a.log.Printf("<Database> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *api) UpdateCustomer(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	id := httptreemux.ContextParams(r.Context())["id"]
	var cus customers.DBCustomer

	GetCustomer, err := a.customer.GetCustomer(id)
	if err != nil {
		a.log.Printf("<Database> %v", err)
		w.WriteHeader(http.StatusNotFound)
	}

	err = json.NewDecoder(r.Body).Decode(&cus)
	if err != nil {
		a.log.Printf("<API> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if cus.FirstName != "" {
		GetCustomer.FirstName = cus.FirstName
	}
	if cus.LastName != "" {
		GetCustomer.LastName = cus.LastName
	}
	if cus.Email != "" {
		GetCustomer.Email = cus.Email
	}
	if cus.Phone != "" {
		GetCustomer.Phone = cus.Phone
	}
	if cus.Address != "" {
		GetCustomer.Address = cus.Address
	}

	err = a.customer.UpdateCustomer(GetCustomer)
	if err != nil {
		a.log.Printf("<Database> %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}