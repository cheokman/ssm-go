package ssm

import "strings"

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

var orgIDMapping = map[string]string{
	"api1op":  "民眾醫療中心（黑沙環）",
	"api2op":  "民眾醫療中心（筷子基）",
	"faom1op": "工人醫療所（台山）",
	"faom2op": "工人醫療所（美的路）",
	"faom4op": "工人醫療所（司打口）",
	"kwerm9":  "澳門鏡湖急診",
	"kwermg":  "澳門鏡湖急診婦科",
	"kwermm":  "澳門鏡湖急診內科",
	"kwermp":  "澳門鏡湖急診兒科",
	"kwermu":  "澳門鏡湖急診未分類",
	"kwermt":  "飛仔鏡湖急診內科",
	"myorgop": "信和醫療中心",
	"ocmoop":  "歸僑總會",
}

// ParsePtsvResponseCode decodes the numeric response code into readable message
func ParseResponseCode(code string) string {
	switch code {
	case "0":
		return "成功 或 符合資格"
	case "2":
		return "沒有提供 求診日期"
	case "3":
		return "沒有提供 證件類別"
	case "4":
		return "沒有提供 證件編號"
	case "12":
		return "求診日期不正確"
	case "13":
		return "求診日期不是上傳當日"
	case "14":
		return "證件類別不正確，必須為 A 至 P 其中一個值"
	case "15":
		return "澳門身份證編號錯誤"
	case "16":
		return "澳門身份證編號不明字元"
	case "21":
		return "不符合資格（需查閱 orgId 與 serviceDate）"
	default:
		return "未知錯誤碼: " + code
	}
}

var ParsePtsvResponseCode = ParsePtsvchkResponseCode

// ParsePtsvchkResponseCode handles 21|orgId|serviceDate format or falls back to code parser
func ParsePtsvchkResponseCode(code string) string {
	if strings.HasPrefix(code, "21|") {
		parts := strings.Split(code, "|")
		if len(parts) == 3 {
			orgId := parts[1]
			serviceDate := parts[2]
			orgName := orgIDMapping[orgId]
			if orgName == "" {
				orgName = "未知機構代碼"
			}
			return "不符合資格（48 小時內重複求診）\n" +
				"上次機構: " + orgId + " - " + orgName + "\n" +
				"上次求診日期: " + serviceDate
		}
		return "不符合資格（返回格式錯誤）: " + code
	}
	return ParseResponseCode(code)
}
