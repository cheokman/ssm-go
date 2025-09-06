package models

type SSMRtssRequest struct {
	WSKey   string
	Action  string // 'u' or 'c'
	FromAM  string
	ToAM    string
	FromPM  string
	ToPM    string
	Waiting string // numeric string, or -1
}

type SSMRtssResponse struct {
	Code    string
	Message string
}
