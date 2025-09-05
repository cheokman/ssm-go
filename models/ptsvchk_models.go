package models

type SSMPtsvChkRequest struct {
	WSKey  string `json:"wskey"`
	Date   string `json:"date"`
	IDType string `json:"idtype"`
	IDNum  string `json:"idnum"`
}

type SSMPtsvChkResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
