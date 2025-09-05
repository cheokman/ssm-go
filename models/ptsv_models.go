package models

type SSMPtsvRequest struct {
	WSKey         string `json:"wskey"`
	Date          string `json:"date"`
	IDType        string `json:"idtype"`
	IDNum         string `json:"idnum"`
	DrLic         string `json:"drliccode"`
	Result1       string `json:"rstcode1"`
	Result2       string `json:"rscode2"`
	Seconds1      string `json:"seconds1"`
	Seconds2      string `json:"seconds2"`
	NoVacReasonID string `json:"novacreasonid"`
}

type SSMPtsvResponse struct {
	Code            string `json:"code"`
	Message         string `json:"message"`
	IsQualified     bool   `json:"is_qualified"`
	OrgId           string `json:"org_id,omitempty"`
	OrgName         string `json:"org_name,omitempty"`
	LastServiceDate string `json:"last_service_date,omitempty"`
}
