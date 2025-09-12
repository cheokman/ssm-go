package rtss

import (
	"net/url"

	"github.com/cheokman/ssm-go/api/ssm"
	"github.com/cheokman/ssm-go/config"
	"github.com/cheokman/ssm-go/internal/logger"
	"github.com/cheokman/ssm-go/internal/response"
	"github.com/cheokman/ssm-go/models"
)

func SubmitWaitingData(req models.SSMRtssRequest, useProd bool) (*models.SSMRtssResponse, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	endpoint := cfg.RTSS.RTSSChkTestURL
	if useProd {
		endpoint = cfg.RTSS.RTSSChkProdURL
	}

	key := cfg.RTSS.WSKey
	if req.WSKey != "" {
		key = req.WSKey
	}

	form := url.Values{}
	form.Set("wskey", key)
	form.Set("action", req.Action)
	form.Set("fromam", req.FromAM)
	form.Set("toam", req.ToAM)
	form.Set("frompm", req.FromPM)
	form.Set("topm", req.ToPM)
	form.Set("waiting", req.Waiting)

	logger.Info("Submitting RTSS data to %s", endpoint)
	logger.Info("Form: %v", form)
	result, err := ssm.DoPostWithRetry(endpoint, form)
	if err != nil {
		return nil, err
	}

	return &models.SSMRtssResponse{
		Code:    result,
		Message: response.ParseRtssResponseCode(result),
	}, nil
}
