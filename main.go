package main

import (
	"context"
	"log"

	"github.com/amarnath-ayyadurai-23/microservices/api"
	"github.com/amarnath-ayyadurai-23/microservices/database"
)

func main(){
	
	ctx := context.Background()
	log := log.Default()
	
	db := database.NewDatabase(ctx, log).GetDB()
	defer db.Close()
	
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	api := api.NewAPI(ctx, db ,log)
	api.Run()
}