// File: models/flu_models.go
package models

type SSMFluRequest struct {
	WSKey         string
	Date          string
	IDType        string
	IDNum         string
	DrLicCode     string
	RstCode1      string
	RstCode2      string
	Seconds1      string
	Seconds2      string
	NoVacReasonID string
}

type SSMFluResponse struct {
	Code    string
	Message string
}
