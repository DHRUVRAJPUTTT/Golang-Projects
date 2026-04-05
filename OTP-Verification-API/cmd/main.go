package main

import (
	"otp-verification-api/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//intialize  config
	app := api.Config{Router: router}
	//routes
	app.Routes()

	router.Run(":8000")
}
