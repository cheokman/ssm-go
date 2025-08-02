package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	WSKey         string `envconfig:"FLU_WS_KEY" required:"true"`
	UseProd       bool   `envconfig:"FLU_USE_PRODUCTION" default:"false"`
	TestURL       string `envconfig:"FLU_TEST_URL" default:"https://www.ssm.gov.mo/outpatient2/flutest.ashx"`
	ProductionURL string `envconfig:"FLU_PROD_URL" default:"https://www.ssm.gov.mo/outpatient2/flu.ashx"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
