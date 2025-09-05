package models

import "encoding/xml"

type SSMPtsvRequest struct {
	XMLName       xml.Name `xml:"ptsv"`
	WSKey         string   `xml:"wskey"`
	Date          string   `xml:"date"`
	IDType        string   `xml:"idtype"`
	IDNum         string   `xml:"idnum"`
	DrLic         string   `xml:"drliccode"`
	Result1       string   `xml:"rstcode1"`
	Result2       string   `xml:"rscode2"`
	Seconds1      string   `xml:"seconds1"`
	Seconds2      string   `xml:"seconds2"`
	NoVacReasonID string   `xml:"novacreasonid"`
}

type SSMPtsvResponse struct {
	XMLName xml.Name `xml:"ptsv_ret"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
}
