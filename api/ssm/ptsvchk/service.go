// File: api/ssm/ptsvchk/service.go
package ptsvchk

import (
	"log"
	"net/url"

	"github.com/cheokman/ssm-go/api/ssm"
	"github.com/cheokman/ssm-go/config"
	"github.com/cheokman/ssm-go/models"
)

// CheckPtsvEligibility 查詢是否符合資格
func CheckPtsvEligibility(req models.SSMPtsvChkRequest, useProd bool) (*models.SSMPtsvChkResponse, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	endpoint := cfg.PTSV.PTSVChkTestURL
	if useProd {
		endpoint = cfg.PTSV.PTSVChkProdURL
	}
	key := cfg.PTSV.WSKey
	if req.WSKey != "" {
		key = req.WSKey
	}
	form := url.Values{}
	form.Set("wskey", key)
	form.Set("date", req.Date)
	form.Set("idtype", req.IDType)
	form.Set("idnum", req.IDNum)

	log.Printf("Checking PTSV eligibility at %s", endpoint)
	result, err := ssm.DoPostWithRetry(endpoint, form)
	if err != nil {
		return nil, err
	}

	parsed := ssm.ParsePtsvchkResponseCode(result)
	return &models.SSMPtsvChkResponse{
		Code:            parsed.Code,
		Message:         parsed.Message,
		IsQualified:     parsed.IsQualified,
		LastServiceDate: parsed.LastServiceDate,
		OrgId:           parsed.OrgID,
		OrgName:         parsed.OrgName,
		OriginalCode:    result,
	}, nil
}
