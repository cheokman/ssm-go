package models

import "encoding/xml"

type SSMPtsvChkRequest struct {
	XMLName   xml.Name `xml:"ptsvchk"`
	WSKey     string   `xml:"wskey"`
	Date      string   `xml:"date"`
	IDType    string   `xml:"idtype"`
	IDNum     string   `xml:"idnum"`
	DrLicCode string   `xml:"drliccode"`
}

type SSMPtsvChkResponse struct {
	XMLName xml.Name `xml:"ptsv_ret"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
}
