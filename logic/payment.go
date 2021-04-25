package logic

import (
	"bookkeeping/dao"
	"bookkeeping/model"
	"fmt"
)

func CreatePayment(p model.PaymentRecord) bool {
	sql_statement := "INSERT INTO paymentrecord (class, date, payment, remark, personID) VALUES ($1, $2, $3, $4, $5);"
	_, err := dao.PostgresDB.Exec(sql_statement, p.Class, p.Date, p.Payment, p.Remark, p.PersonID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
