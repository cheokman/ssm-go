package flu

import (
	"log"
	"net/url"

	"github.com/cheokman/ssm-go/api/ssm"
	"github.com/cheokman/ssm-go/models"
)

const (
	testURL       = "https://www.ssm.gov.mo/outpatient2/flutest.ashx"
	productionURL = "https://www.ssm.gov.mo/outpatient2/flu.ashx"
)

func SendFluVaccineData(req models.SSMFluRequest, useProd bool) (*models.SSMFluResponse, error) {
	endpoint := testURL
	if useProd {
		endpoint = productionURL
	}

	form := url.Values{}
	form.Set("wskey", req.WSKey)
	form.Set("date", req.Date)
	form.Set("idtype", req.IDType)
	form.Set("idnum", req.IDNum)
	form.Set("drliccode", req.DrLicCode)
	form.Set("rstcode1", req.RstCode1)
	form.Set("rstcode2", req.RstCode2)
	form.Set("seconds1", req.Seconds1)
	form.Set("seconds2", req.Seconds2)
	form.Set("novacreasonid", req.NoVacReasonID)

	log.Printf("Sending Flu Vaccine Data to %s", endpoint)
	result, err := ssm.DoPostWithRetry(endpoint, form)
	if err != nil {
		return nil, err
	}

	return &models.SSMFluResponse{
		Code:    result,
		Message: ssm.ParseFluResponseCode(result),
	}, nil
}
