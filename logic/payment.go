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

func GetPieData(userID string) ([]model.PeiData, bool) {
	tranlateMap := map[string]string{
		"eat":           "食",
		"clothes":       "衣",
		"live":          "住",
		"traffic":       "行",
		"educate":       "育",
		"entertainment": "樂",
	}
	sql_statement := `SELECT class, SUM(payment) FROM paymentrecord WHERE personid=$1 GROUP BY class`
	rows, err := dao.PostgresDB.Query(sql_statement, userID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	var pieRet []model.PeiData
	var tmpData model.PeiData
	for rows.Next() {
		if err := rows.Scan(&tmpData.Class, &tmpData.Total); err != nil {
			fmt.Println(err.Error())
			return nil, false
		}
		tmpData.Class = tranlateMap[tmpData.Class]
		pieRet = append(pieRet, tmpData)
	}
	return pieRet, true
}

func GetLineChartData(userID string) ([]model.LineData, bool) {
	tranlateMap := map[string]string{
		"1":  "1月",
		"2":  "2月",
		"3":  "3月",
		"4":  "4月",
		"5":  "5月",
		"6":  "6月",
		"7":  "7月",
		"8":  "8月",
		"9":  "9月",
		"10": "10月",
		"11": "11月",
		"12": "12月",
	}
	sql_statement := `SELECT DATE_PART('MONTH', date), SUM(payment) TotalCount FROM paymentrecord WHERE date>=$1 AND date<=$2 AND personid=$3 GROUP BY DATE_PART('MONTH', date);`
	year := strconv.Itoa(time.Now().Year())
	yearlimitStart := year + `-01-01`
	yearlimitEnd := year + `-12-31`
	rows, err := dao.PostgresDB.Query(sql_statement, yearlimitStart, yearlimitEnd, userID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	var lineRet []model.LineData
	var tmpData model.LineData
	for rows.Next() {
		if err := rows.Scan(&tmpData.Month, &tmpData.Total); err != nil {
			fmt.Println(err.Error())
			return nil, false
		}
		tmpData.Month = tranlateMap[tmpData.Month]
		lineRet = append(lineRet, tmpData)
	}
	return lineRet, true
}
