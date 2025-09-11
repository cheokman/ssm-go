# SSM Outpatient Integration Configurations

This package provides configuration structs for integrating with **SSM (Serviços de Saúde de Macau)** web services, including Flu Vaccination, PTSV, and RTSS endpoints.  
Configuration is loaded using `envconfig` with environment variable support and sensible defaults.

---

## FluConfig

Configuration for Flu Vaccination Service.

| Field           | Env Var              | Type   | Default                                           | Required | Description                    |
| --------------- | -------------------- | ------ | ------------------------------------------------- | -------- | ------------------------------ |
| `WSKey`         | `FLU_WS_KEY`         | string | —                                                 | false    | Web service key (if required). |
| `UseProd`       | `FLU_USE_PRODUCTION` | bool   | `false`                                           | false    | Switch between test and prod.  |
| `TestURL`       | `FLU_TEST_URL`       | string | `https://www.ssm.gov.mo/outpatient2/flutest.ashx` | false    | Endpoint for test environment. |
| `ProductionURL` | `FLU_PROD_URL`       | string | `https://www.ssm.gov.mo/outpatient2/flu.ashx`     | false    | Endpoint for production.       |

---

## PTSVConfig

Configuration for PTSV Service and its check endpoints.

| Field            | Env Var               | Type   | Default                                               | Required | Description                     |
| ---------------- | --------------------- | ------ | ----------------------------------------------------- | -------- | ------------------------------- |
| `WSKey`          | `PTSV_WS_KEY`         | string | —                                                     | false    | Web service key (if required).  |
| `UseProduction`  | `PTSV_USE_PRODUCTION` | bool   | `false`                                               | false    | Switch between test and prod.   |
| `PTSVTestURL`    | `PTSV_TEST_URL`       | string | `https://www.ssm.gov.mo/outpatient2/ptsvtest.ashx`    | false    | PTSV test environment endpoint. |
| `PTSVProdURL`    | `PTSV_PROD_URL`       | string | `https://www.ssm.gov.mo/outpatient2/ptsv.ashx`        | false    | PTSV production endpoint.       |
| `PTSVChkTestURL` | `PTSVCHK_TEST_URL`    | string | `https://www.ssm.gov.mo/outpatient2/ptsvchktest.ashx` | false    | PTSV check (test).              |
| `PTSVChkProdURL` | `PTSVCHK_PROD_URL`    | string | `https://www.ssm.gov.mo/outpatient2/ptsvchk.ashx`     | false    | PTSV check (production).        |

---

## RTSSConfig

Configuration for RTSS Check Service.

| Field            | Env Var            | Type   | Default                                            | Required | Description                     |
| ---------------- | ------------------ | ------ | -------------------------------------------------- | -------- | ------------------------------- |
| `WSKey`          | `RTSS_WS_KEY`      | string | —                                                  | false    | Web service key (if required).  |
| `RTSSChkTestURL` | `RTSSCHK_TEST_URL` | string | `https://www.ssm.gov.mo/outpatient2/rtsstest.ashx` | false    | RTSS test check endpoint.       |
| `RTSSChkProdURL` | `RTSSCHK_PROD_URL` | string | `https://www.ssm.gov.mo/outpatient2/rtss.ashx`     | false    | RTSS production check endpoint. |

---

## Usage

```go
package main

import (
    "fmt"
    "github.com/kelseyhightower/envconfig"
)

func main() {
    var flu FluConfig
    var ptsv PTSVConfig
    var rtss RTSSConfig

    envconfig.Process("", &flu)
    envconfig.Process("", &ptsv)
    envconfig.Process("", &rtss)

    fmt.Printf("Flu Config: %+v\n", flu)
    fmt.Printf("PTSV Config: %+v\n", ptsv)
    fmt.Printf("RTSS Config: %+v\n", rtss)
}
```
