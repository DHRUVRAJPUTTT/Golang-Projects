package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Add this struct so the rest of your app can use it!
type Config struct {
	Router *gin.Engine
}

func envACCOUNTSID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

// ... keep envAUTHTOKEN and envSERVICESID exactly as they were
func envAUTHTOKEN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}
	return os.Getenv(("TWILIO_AUTHTOKEN"))
}
func envSERVICESID() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}
	return os.Getenv("TWILIO_SERVICES_ID")

}
