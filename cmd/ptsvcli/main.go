package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cheokman/ssm-go/api/ssm/ptsv"
	"github.com/cheokman/ssm-go/models"
)

func main() {
	var (
		useProd       bool
		date          string
		idtype        string
		idnum         string
		drliccode     string
		rstcode1      string
		rscode2       string
		seconds1      string
		seconds2      string
		novacreasonid string
		wskeyOverride string
	)

	flag.BoolVar(&useProd, "prod", false, "Use production endpoint")
	flag.StringVar(&date, "date", "", "求診日期 YYYYMMDD")
	flag.StringVar(&idtype, "idtype", "P", "證件類別（預設 P）")
	flag.StringVar(&idnum, "idnum", "", "證件號碼")
	flag.StringVar(&drliccode, "drlic", "", "醫生牌照號碼")
	flag.StringVar(&rstcode1, "r1", "", "視窗1選擇結果 (Y/N)")
	flag.StringVar(&rscode2, "r2", "", "視窗2選擇結果 (S/U/N)")
	flag.StringVar(&seconds1, "s1", "", "秒數1")
	flag.StringVar(&seconds2, "s2", "", "秒數2")
	flag.StringVar(&novacreasonid, "reason", "0000000000", "未接種原因代碼（預設全0）")
	flag.StringVar(&wskeyOverride, "wskey", "", "自訂 wskey（可選）")

	flag.Parse()

	req := models.SSMPtsvRequest{
		WSKey:         wskeyOverride, // will be overridden by config internally if empty
		Date:          date,
		IDType:        idtype,
		IDNum:         idnum,
		DrLic:         drliccode,
		Result1:       rstcode1,
		Result2:       rscode2,
		Seconds1:      seconds1,
		Seconds2:      seconds2,
		NoVacReasonID: novacreasonid,
	}

	resp, err := ptsv.SendPtsvData(req, useProd)
	if err != nil {
		log.Fatalf("送出失敗: %v", err)
	}

	fmt.Println("=== 提交結果 ===")
	fmt.Printf("原始代碼: %s\n", resp.Code)
	fmt.Printf("說明訊息: %s\n", resp.Message)
	if !resp.IsQualified {
		fmt.Printf("上次求診日期: %s\n", resp.LastServiceDate)
		fmt.Printf("機構代號: %s\n", resp.OrgId)
		fmt.Printf("機構名稱: %s\n", resp.OrgName)
	}
}
