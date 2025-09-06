package response

func ParseRtssResponseCode(code string) string {
	switch code {
	case "0":
		return "提交成功"
	case "17":
		return "候診人數錯誤（非正整數或 -1）"
	default:
		return "未知錯誤碼: " + code
	}
}
