package response

func ParseFluResponseCode(code string) string {
	switch code {
	case "0":
		return "儲存成功"
	case "2":
		return "沒有提供求診日期"
	case "3":
		return "沒有提供證件類別"
	case "4":
		return "沒有提供證件編號"
	case "5":
		return "沒有提供醫生牌照號碼"
	case "6":
		return "沒有提供視窗 1 選擇結果"
	case "7":
		return "沒有提供視窗 2 選擇結果"
	case "10":
		return "沒有提供未接種原因"
	case "21":
		return "求診日期格式錯誤"
	case "31":
		return "證件類別無效"
	default:
		return "未知錯誤碼: " + code
	}
}
