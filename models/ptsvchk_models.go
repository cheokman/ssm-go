package models

type SSMPtsvChkRequest struct {
	WSKey  string `json:"wskey"`
	Date   string `json:"date"`
	IDType string `json:"idtype"`
	IDNum  string `json:"idnum"`
}

type SSMPtsvChkResponse struct {
	Code            string `json:"code"`
	Message         string `json:"message"`
	IsQualified     bool   `json:"is_qualified"`
	OrgId           string `json:"org_id,omitempty"`
	OrgName         string `json:"org_name,omitempty"`
	LastServiceDate string `json:"last_service_date,omitempty"`
}
