package controller

import (
	"bookkeeping/config"
	"bookkeeping/dao"
	"bookkeeping/logic"
	"bookkeeping/model"
	"fmt"
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

func GetPaymentHistory(c *gin.Context) {
	userID := c.MustGet(config.AuthMidUserNameKey).(string)

	sql_statement := `SELECT class, payment, date, remark FROM paymentrecord WHERE personid=$1`
	rows, err := dao.PostgresDB.Query(sql_statement, userID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	retData := []model.PaymentRecord{}
	var tmpData model.PaymentRecord

	for rows.Next() {
		if err := rows.Scan(&tmpData.Class, &tmpData.Payment, &tmpData.Date, &tmpData.Remark); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, nil)
		}
		retData = append(retData, tmpData)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": retData,
	})
}
