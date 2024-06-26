package customers

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/amarnath-ayyadurai-23/microservices/database"
<<<<<<< HEAD
	"github.com/google/uuid"

=======

	"github.com/google/uuid"
>>>>>>> ddcc85d (github actions update)
	"github.com/stretchr/testify/assert"
)

var (
	dbase	*database.Database
	customer *Customer
)

func init(){
	ctx := context.Background()
	log := log.New(os.Stdout, "test", 1)
<<<<<<< HEAD
	//  // Mock database connection
	//  dbase = database.NewDatabase(ctx, log)
	//  customer = NewCustomer(ctx, dbase.GetDB(), log)
=======
	 // Mock database connection
	 dbase = database.NewDatabase(ctx, log)
	 customer = NewCustomer(ctx, dbase.GetDB(), log)
>>>>>>> ddcc85d (github actions update)
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
	
	//Call GetCustomer function
	customers, err := customer.QuerybyEmail("adipiscing.elit.Etiam@euultricessit.edu")

	// Assert no error occurred
	assert.Nil(t, err)

	// Assert customer ID is correct
	assert.Equal(t,customers[0].FirstName,"Dorian")
	assert.Equal(t,customers[0].Email,"adipiscing.elit.Etiam@euultricessit.edu")

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