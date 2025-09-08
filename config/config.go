package config

import "github.com/kelseyhightower/envconfig"

// type Config struct {
// 	WSKey         string `envconfig:"FLU_WS_KEY" required:"true"`
// 	UseProd       bool   `envconfig:"FLU_USE_PRODUCTION" default:"false"`
// 	TestURL       string `envconfig:"FLU_TEST_URL" default:"https://www.ssm.gov.mo/outpatient2/flutest.ashx"`
// 	ProductionURL string `envconfig:"FLU_PROD_URL" default:"https://www.ssm.gov.mo/outpatient2/flu.ashx"`
// }

type FluConfig struct {
	WSKey         string `envconfig:"FLU_WS_KEY" required:"false"`
	UseProd       bool   `envconfig:"FLU_USE_PRODUCTION" default:"false"`
	TestURL       string `envconfig:"FLU_TEST_URL" default:"https://www.ssm.gov.mo/outpatient2/flutest.ashx"`
	ProductionURL string `envconfig:"FLU_PROD_URL" default:"https://www.ssm.gov.mo/outpatient2/flu.ashx"`
}

type PTSVConfig struct {
	WSKey          string `envconfig:"PTSV_WS_KEY" required:"false"`
	UseProduction  bool   `envconfig:"PTSV_USE_PRODUCTION" default:"false"`
	PTSVTestURL    string `envconfig:"PTSV_TEST_URL" default:"https://www.ssm.gov.mo/outpatient2/ptsvtest.ashx"`
	PTSVProdURL    string `envconfig:"PTSV_PROD_URL" default:"https://www.ssm.gov.mo/outpatient2/ptsv.ashx"`
	PTSVChkTestURL string `envconfig:"PTSVCHK_TEST_URL" default:"https://www.ssm.gov.mo/outpatient2/ptsvchktest.ashx"`
	PTSVChkProdURL string `envconfig:"PTSVCHK_PROD_URL" default:"https://www.ssm.gov.mo/outpatient2/ptsvchk.ashx"`
}

type RTSSConfig struct {
	WSKey          string `envconfig:"RTSS_WS_KEY" required:"false"`
	RTSSChkTestURL string `envconfig:"RTSSCHK_TEST_URL" default:"https://www.ssm.gov.mo/outpatient2/rtsstest.ashx"`
	RTSSChkProdURL string `envconfig:"RTSSCHK_PROD_URL" default:"https://www.ssm.gov.mo/outpatient2/rtss.ashx"`
}

type Config struct {
	Flu  FluConfig
	PTSV PTSVConfig
	RTSS RTSSConfig
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
