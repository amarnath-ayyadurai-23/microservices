package customers

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/amarnath-ayyadurai-23/microservices/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// var (
// 	dbase	*database.Database
// 	customer *Customer
// )

func init(){
	ctx := context.Background()
	log := log.New(os.Stdout, "test", 1)
	 // Mock database connection
	 dbase = database.NewDatabase(ctx, log)
	 customer = NewCustomer(ctx, dbase.GetDB(), log)
}

func Test_GetCustomer(t *testing.T) {
   
    // Call GetCustomer function
    customers, err := customer.GetCustomers()

    // Assert no error occurred
    assert.Nil(t, err)

    // Assert customer ID is correct
    assert.Greater(t, len(customers),10)
}

func Test_GetCustomerByID(t *testing.T) {	
	
	// Call GetCustomer function
	customers, err := customer.GetCustomer("653c97a6-044c-42e6-a698-f65d6432dfd2")

	// // Assert no error occurred
	// assert.Nil(t, err)

	// Assert customer ID is correct
	assert.Equal(t,customers.FirstName,"Dorian")
	assert.Equal(t,customers.Email,"adipiscing.elit.Etiam@euultricessit.edu")

}

func Test_createUpdateNDelete(t *testing.T) {	
	
	id := uuid.New().String() 
	// Create customer
	cus:= DBCustomer{
		ID: id,
		FirstName: "Test",
		LastName: "User",
		Email: "test@test.com",
	}

	// Call CreateCustomer function
	err := customer.CreateCustomer(cus)
	assert.Nil(t, err)
	assert.Equal(t, cus.Phone, "")
	assert.Equal(t, cus.Address, "")

	// update customer
	cus.Phone = "123-456-7890"
	cus.Address = "123 Main St"
	err = customer.UpdateCustomer(cus)
	assert.Nil(t, err)

	// delete customer
	err = customer.DeletebyID(id)
	assert.Nil(t, err)

}