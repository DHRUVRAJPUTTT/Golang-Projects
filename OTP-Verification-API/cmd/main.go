package main

import (
	"github.com/DHRUVRAJPUTTT/Golang-Projects/tree/main/OTP-Verification-API/api"
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
