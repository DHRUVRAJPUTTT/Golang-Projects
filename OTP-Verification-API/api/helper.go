package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	// Changed BlindJSON to ShouldBindJSON
	if err := c.ShouldBindJSON(&data); err != nil {
		return err
	}
	return nil
}

// 19.01
func (app *Config) writeJSON(c *gin.Context, status int, data any) {
	c.JSON(status, jsonResponse{Status: status, Message: "success", Data: data})
}
func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
}
