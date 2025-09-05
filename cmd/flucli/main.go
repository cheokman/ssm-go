// File: cmd/flucli/main.go
package main

import (
	"log"
	"os"

	"github.com/cheokman/ssm-go/api/ssm/flu"
	"github.com/cheokman/ssm-go/config"
	"github.com/cheokman/ssm-go/models"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	log.SetOutput(os.Stdout)
	log.Println("\u1f680 Sending Flu Vaccine Data")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}
	log.Printf("WSKey: %s \n", cfg.Flu.WSKey)
	req := models.SSMFluRequest{
		WSKey:         cfg.Flu.WSKey,
		Date:          "20250808",
		IDType:        "P",
		IDNum:         "12345678",
		DrLicCode:     "MI0001",
		RstCode1:      "Y",
		RstCode2:      "S",
		Seconds1:      "3",
		Seconds2:      "15",
		NoVacReasonID: "0000000000",
	}

	resp, err := flu.SendFluVaccineData(req, cfg.Flu.UseProd)
	if err != nil {
		log.Fatalf("send failed: %v", err)
	}
	log.Printf("\u2705 Result: %s (%s)", resp.Code, resp.Message)
}
