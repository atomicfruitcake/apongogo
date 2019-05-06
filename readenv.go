package readEnv

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

var mongoUser, mongoPass, mongoUrl string

func readEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	mongoUser = os.Getenv("MONGODB_USER")
	mongoPass = os.Getenv("MONGODB_PASS")
	mongoUrl =  os.Getenv("MONGODB_URL")
}	