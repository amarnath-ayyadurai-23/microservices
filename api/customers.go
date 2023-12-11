package api

import (
	"encoding/json"
	"net/http"
	"unsafe"

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