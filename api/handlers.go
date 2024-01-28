package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sriganeshres/creditcardverifier/data"
)

const appTimeOut = time.Second * 10

var times = 0

func (app *Config) VerifyCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeOut)
		var payload data.CardData
		times++
		fmt.Println("Requested ", times)
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.CardData{
			CardNumber: payload.CardNumber,
		}

		val := app.valiadateLuhn(newData.CardNumber)
		if val {
			app.writeJSON(c, http.StatusAccepted, "Card is Valid", "Success")
		} else {
			app.writeJSON(c, http.StatusBadRequest, "Card is inValid", "Failure")
		}
	}
}
