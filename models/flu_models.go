// File: models/flu_models.go
package models

type SSMFluRequest struct {
	WSKey         string `json:"wskey"`
	Date          string `json:"date"`
	IDType        string `json:"idtype"`
	IDNum         string `json:"idnum"`
	DrLicCode     string `json:"drliccode"`
	RstCode1      string `json:"rstcode1"`
	RstCode2      string `json:"rstcode2"`
	Seconds1      string `json:"seconds1"`
	Seconds2      string `json:"seconds2"`
	NoVacReasonID string `json:"novacreasonid"`
}

type SSMFluResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
