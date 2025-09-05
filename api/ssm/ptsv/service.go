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
	form.Set("wskey", req.WSKey)
	form.Set("date", req.Date)
	form.Set("idtype", req.IDType)
	form.Set("idnum", req.IDNum)
	form.Set("drliccode", req.DrLic)
	form.Set("rstcode1", req.Result1)
	form.Set("rscode2", req.Result2)
	form.Set("seconds1", req.Seconds1)
	form.Set("seconds2", req.Seconds2)
	form.Set("novacreasonid", req.NoVacReasonID)

	log.Printf("Sending PTSV data to %s", endpoint)
	result, err := ssm.DoPostWithRetry(endpoint, form)
	if err != nil {
		return nil, err
	}

	return &models.SSMPtsvResponse{
		Code:    result,
		Message: ssm.ParsePtsvResponseCode(result),
	}, nil
}
