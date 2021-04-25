package model

type PaymentRecord struct {
	ID       string `json:"id" form:"id"`
	Class    string `json:"class" form:"class"`
	Payment  string `json:"payment" form:"payment"`
	Date     string `json:"date" form:"date"`
	Remark   string `json:"remark" form:"remark"`
	PersonID string `json:"personID" form:"personID"`
}
