package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cheokman/ssm-go/api/ssm/ptsvchk"
	"github.com/cheokman/ssm-go/models"
)

func main() {
	var (
		useProd       bool
		date          string
		idtype        string
		idnum         string
		drliccode     string
		wskeyOverride string
	)

	flag.BoolVar(&useProd, "prod", false, "Use production endpoint")
	flag.StringVar(&date, "date", "", "求診日期 YYYYMMDD")
	flag.StringVar(&idtype, "idtype", "P", "證件類別（預設 P）")
	flag.StringVar(&idnum, "idnum", "", "證件號碼")
	flag.StringVar(&drliccode, "drlic", "", "醫生牌照號碼")
	flag.StringVar(&wskeyOverride, "wskey", "", "自訂 wskey（可選）")

	flag.Parse()

	req := models.SSMPtsvChkRequest{
		WSKey:     wskeyOverride, // will be overridden by config if empty
		Date:      date,
		IDType:    idtype,
		IDNum:     idnum,
		DrLicCode: drliccode,
	}

	resp, err := ptsvchk.CheckPtsvEligibility(req, useProd)
	if err != nil {
		log.Fatalf("查詢失敗: %v", err)
	}

	fmt.Println("=== 查詢結果 ===")
	fmt.Printf("原始代碼: %s\n", resp.Code)
	fmt.Printf("說明訊息: %s\n", resp.Message)
}
