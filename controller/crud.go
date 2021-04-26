package controller

import (
	"bookkeeping/config"
	"bookkeeping/dao"
	"bookkeeping/logic"
	"bookkeeping/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddPayment(c *gin.Context) {
	userID := c.MustGet(config.AuthMidUserNameKey).(string)
	var input model.PaymentRecord
	c.Bind(&input)
	if input.Class == "" || input.Date == "" || input.Payment == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	input.PersonID = userID
	if logic.CreatePayment(input) == false {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func GetPaymentHistory(c *gin.Context) {
	userID := c.MustGet(config.AuthMidUserNameKey).(string)
	tranlateMap := map[string]string{
		"eat":           "食",
		"clothes":       "衣",
		"live":          "住",
		"traffic":       "行",
		"educate":       "育",
		"entertainment": "樂",
	}
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
		tmpData.Date = strings.ReplaceAll(tmpData.Date, "T00:00:00Z", "")
		tmpData.Class = tranlateMap[tmpData.Class]
		retData = append(retData, tmpData)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": retData,
	})
}
