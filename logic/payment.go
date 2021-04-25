package logic

import (
	"bookkeeping/dao"
	"bookkeeping/model"
	"fmt"
	"strconv"
	"time"
)

func CreatePayment(p model.PaymentRecord) bool {
	sql_statement := "INSERT INTO paymentrecord (class, date, payment, remark, personID) VALUES ($1, $2, $3, $4, $5);"
	payment, _ := strconv.Atoi(p.Payment)
	date, _ := time.Parse("2006-01-02", p.Date)
	_, err := dao.PostgresDB.Exec(sql_statement, p.Class, date, payment, p.Remark, p.PersonID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
