package model

type PeiData struct {
	Class string `json:"class" form:"class"`
	Total int    `json:"total" form:"total"`
}

type LineData struct {
	Month string `json:"month" form:"month"`
	Total int    `json:"total" form:"total"`
}
