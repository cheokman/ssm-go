package main

import (
	"fmt"
	"os"

	"github.com/cheokman/ssm-go/api/ssm/rtss"
	"github.com/cheokman/ssm-go/config"
	"github.com/cheokman/ssm-go/models"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	req := models.SSMRtssRequest{
		WSKey:   cfg.PTSV.WSKey,
		Action:  "u", // or "c"
		FromAM:  "09:00",
		ToAM:    "13:00",
		FromPM:  "16:00",
		ToPM:    "20:00",
		Waiting: "12", // or "-1"
	}

	resp, err := rtss.SubmitWaitingData(req, cfg.RTSS.UseProd)
	if err != nil {
		fmt.Println("提交失敗:", err)
		os.Exit(1)
	}

	fmt.Printf("回傳碼: %s\n訊息: %s\n", resp.Code, resp.Message)
}
