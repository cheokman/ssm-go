package middleware

import (
	"io"
	"net/http"
	"time"

	"github.com/cheokman/ssm-go/internal/logger"
	"github.com/hashicorp/go-retryablehttp"
)

var client = retryablehttp.NewClient()

func init() {
	client.RetryMax = 3
	client.RetryWaitMin = 500 * time.Millisecond
	client.RetryWaitMax = 3 * time.Second
	client.Logger = nil // 關閉預設 stdout logging
}

// PostWithRetry sends an HTTP POST request with automatic retry and logging
func PostWithRetry(url string, body io.Reader, contentType string) (*http.Response, error) {
	req, err := retryablehttp.NewRequest("POST", url, body)
	if err != nil {
		logger.Error("Failed to create request: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	logger.Info("Sending POST to %s", url)
	return client.Do(req)
}
