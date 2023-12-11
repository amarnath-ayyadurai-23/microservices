package customers

import (
	"context"
	"fmt"
	"log"

	"github.com/amarnath-ayyadurai-23/microservices/database"

	"github.com/jmoiron/sqlx"
)

type DBCustomer struct {
	ID        string `db:"customer_id" json:"customer_id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	Phone     string `db:"phone" json:"phone"`
	Address   string `db:"address" json:"address"`
}

type Customer struct {
	ctx context.Context
	db  *sqlx.DB
	log *log.Logger
}

func NewCustomer(ctx context.Context, db *sqlx.DB, log *log.Logger) *Customer {
	return &Customer{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (c *Customer) LogF(format string) string {
	return fmt.Sprintf("[Customer] %s", format)
}

func (c *Customer) GetCustomers() ([]DBCustomer, error){
	
	var customers []DBCustomer
	query := `SELECT * FROM wisdom.customers`
	data := struct {}{}

	err := database.NamedQuerySlice(c.ctx, c.log, c.db, query, data, &customers)
	if err != nil {
		c.log.Printf("<Customer> %v", err)
		return []DBCustomer{}, err
	}
	// rows, err := sqlx.NamedQueryContext(c.ctx, c.db, query, struct{}{})
	// if err != nil {
	// 	c.log.Printf("<Customer> %v", err)
	// 	return []DBCustomer{}, err
	// }
	// defer rows.Close() //nolint:all

	// for rows.Next() {
	// 	var cus DBCustomer
	// 	err := rows.StructScan(&cus)
	// 	if err != nil {
	// 		c.log.Printf("<Customer> %v",err)
	// 		return []DBCustomer{}, err
	// 	}
	// 	customers = append(customers, cus)
	// }
	c.log.Println(c.LogF("Getting customers"))
	return customers, nil
}

func (c *Customer) GetCustomer(id string) (DBCustomer, error){
	var customer DBCustomer
	query := `SELECT * FROM wisdom.customers WHERE customer_id = :customer_id`
	data := struct {
		ID string `db:"customer_id"`
	}{ID: id}
	err := database.NamedQueryStruct(c.ctx, c.log, c.db, query, data, &customer)
	if err != nil {
		c.log.Printf("<Customer> %v", err)
		return DBCustomer{}, err
	}

	return customer,nil
}
