package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cheokman/ssm-go/api/ssm/ptsvchk"
	"github.com/cheokman/ssm-go/config"
	"github.com/cheokman/ssm-go/models"
)

func main() {
	var (
		useProd bool
		date    string
		idtype  string
		idnum   string
	)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}
	log.Printf("WSKey: %s \n", cfg.Flu.WSKey)

	flag.BoolVar(&useProd, "prod", false, "Use production endpoint")
	flag.StringVar(&date, "date", "", "求診日期 YYYYMMDD")
	flag.StringVar(&idtype, "idtype", "P", "證件類別（預設 P）")
	flag.StringVar(&idnum, "idnum", "", "證件號碼")

	flag.Parse()

	req := models.SSMPtsvChkRequest{
		WSKey:  cfg.PTSV.WSKey, // will be overridden by config if empty
		Date:   date,
		IDType: idtype,
		IDNum:  idnum,
	}

	resp, err := ptsvchk.CheckPtsvEligibility(req, useProd)
	if err != nil {
		log.Fatalf("查詢失敗: %v", err)
	}

	fmt.Println("=== 查詢結果 ===")
	fmt.Printf("原始代碼: %s\n", resp.Code)
	fmt.Printf("說明訊息: %s\n", resp.Message)
	if !resp.IsQualified {
		fmt.Printf("上次求診日期: %s\n", resp.LastServiceDate)
		fmt.Printf("機構代號: %s\n", resp.OrgId)
		fmt.Printf("機構名稱: %s\n", resp.OrgName)
	}
}
