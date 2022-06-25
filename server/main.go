package main

import (
	"os"
	// "log"

	"github.com/jessetuglu/bill_app/server/api"
	"github.com/jessetuglu/bill_app/server/db"
	// "github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func main(){
	var err error

	oauth_client_config := oauth2.Config{
		Endpoint: google.Endpoint,
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL: os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{"openid", "email", "profile"},
	}

	base, _ := zap.NewProduction()
	logger := base.Sugar()
	db, err := db.ConnectToDB(os.Getenv("BILL_DB_USER"), os.Getenv("BILL_DB_PASSWORD"), os.Getenv("BILL_DB_URL"), os.Getenv("BILL_DB_DB_NAME"))
	if (err != nil){
		logger.Fatalw("Error: Couldn't connect to DB", err)
		return
	}

	server := api.NewServer(logger, db, oauth_client_config)
	server.Router.Run(":"+os.Getenv("SERVER_PORT"))
}