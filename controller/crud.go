package controller

import (
	"bookkeeping/config"
	"bookkeeping/logic"
	"bookkeeping/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPayment(c *gin.Context) {
	userID := c.MustGet(config.AuthMidUserNameKey).(string)
	var input model.PaymentRecord
	c.Bind(&input)
	input.PersonID = userID
	if logic.CreatePayment(input) == false {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
