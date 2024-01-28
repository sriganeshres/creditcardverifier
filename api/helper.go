package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var valiadate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	if err := valiadate.Struct(&data); err != nil {
		return err
	}
	return nil
}

func (app *Config) writeJSON(c *gin.Context, status int, data any, message string) {
	c.JSON(status, jsonResponse{Status: status, Message: message, Data: data})
}

func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
}

func (app *Config) valiadateLuhn(cardNum string) bool {
	checkSum, err := strconv.Atoi(string(cardNum[len(cardNum)-1]))
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	sum := 0
	for i := len(cardNum) - 2; i >= 0; i-- {
		val := 0
		digit, e := strconv.Atoi(string(cardNum[i]))
		if e != nil {
			fmt.Println("Error: ", e)
			return false
		}
		if i%2 == 0 {
			digit *= 2
		}
		val = digit/10 + digit%10
		sum += val
	}

	return 10-sum%10 == checkSum
}
