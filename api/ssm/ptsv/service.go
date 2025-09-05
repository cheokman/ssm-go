package ptsv

import (
	"log"
	"net/url"

	"github.com/cheokman/ssm-go/api/ssm"
	"github.com/cheokman/ssm-go/config"
	"github.com/cheokman/ssm-go/models"
)

// SendPtsvData sends outpatient vaccine data to SSM (ptsv endpoint)
func SendPtsvData(req models.SSMPtsvRequest, useProd bool) (*models.SSMPtsvResponse, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	endpoint := cfg.PTSV.PTSVTestURL
	if useProd {
		endpoint = cfg.PTSV.PTSVProdURL
	}

	form := url.Values{}
	form.Set("wskey", cfg.PTSV.WSKey)
	form.Set("date", req.Date)
	form.Set("idtype", req.IDType)
	form.Set("idnum", req.IDNum)

	log.Printf("Checking PTSV eligibility at %s", endpoint)
	result, err := ssm.DoPostWithRetry(endpoint, form)
	if err != nil {
		return nil, err
	}
	parsed := ssm.ParsePtsvResponseCode(result)
	return &models.SSMPtsvResponse{
		Code:            result,
		Message:         parsed.Message,
		IsQualified:     parsed.IsQualified,
		LastServiceDate: parsed.LastServiceDate,
		OrgId:           parsed.OrgID,
		OrgName:         parsed.OrgName,
	}, nil
}
