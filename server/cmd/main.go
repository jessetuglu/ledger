package main

import (
	"log"
	"github.com/jessetuglu/bill_app/server"
	"github.com/joho/godotenv"
)


func main(){
	var err error
	if err = godotenv.Load("server/config/prod.env"); err != nil{
		log.Fatal("Error: couldn't load env variables")
	}
	svr := server.InitializeServer()
	
	
	// go svr.Router.ServeHTTP()
}