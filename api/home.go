package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amarnath-ayyadurai-23/microservices/models/customers"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/jmoiron/sqlx"
)

type api struct {
	ctx      context.Context
	db       *sqlx.DB
	router   *httptreemux.ContextMux
	log      *log.Logger
	customer *customers.Customer
}

func NewAPI(ctx context.Context, db *sqlx.DB, log *log.Logger) *api {
	return &api{
		ctx:      ctx,
		db:       db,
		router:   httptreemux.NewContextMux(),
		log:      log,
		customer: customers.NewCustomer(ctx, db, log),
	}
}

func (a *api) LogF(format string) string {
	return fmt.Sprintf("[API] %s", format)
}


func (a *api) Run() {

	a.router.Handle("GET", "/",a.Ping)
	
	a.router.Handle("GET", "/customers", a.GetCustomers)
	a.router.Handle("GET", "/customers/:id", a.GetCustomer)

	a.log.Println(a.LogF("Starting server on port 8080"))
	http.ListenAndServe(":8080", a.router)

}

func (a *api) Ping(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, time.Now().String())
	a.log.Println(a.LogF("Pinging server on port 8080"))
}