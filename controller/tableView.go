package controller

import (
	"fmt"
	"time"
)

func TableView() {

	var timess int64 = 1639444750958
	tm := time.UnixMilli(timess)
	// fmt.Println(tm)

	// y := tm.Year()
	// m := tm.Month()
	// d := tm.Day()
	// h := tm.Hour()
	// i := tm.Minute()
	// s := tm.Second()
	// fmt.Println(y, m, d, h, i, s)
	// data_nCov.Results[0].Cities[localIp_city].CurrentConfirmedCount
	fmt.Println("时间", "城市", "现存确诊", "累计确诊", "可疑计数", "治愈数", "死亡人数")
	for i := 0; i < len(data_nCov.Results[0].Cities); i++ {
		fmt.Println(tm, data_nCov.Results[0].Cities[i].CityName, data_nCov.Results[0].Cities[i].CurrentConfirmedCount, data_nCov.Results[0].Cities[i].ConfirmedCount, data_nCov.Results[0].Cities[i].SuspectedCount, data_nCov.Results[0].Cities[i].CuredCount, data_nCov.Results[0].Cities[i].DeadCount)
	}
	// data1 := [][]string{
	// 	{"12/14", "", "", "d", "d", "d", "d"},
	// 	{"12/14", "", "", "d", "d", "d", "d"},
	// 	{"12/14", "dd", "dd", "s", "h", "c", "s"},
	// }

	// tableList := tablewriter.NewWriter(os.Stdout)
	// tableList.SetHeader([]string{"时间", "所在地", "现存确诊", "累计确诊", "可疑计数", "治愈数", "死亡人数"})
	// tableList.SetFooter([]string{"", "", "sd", "", "sd", "", ""})
	// tableList.SetBorder(false)

	// tableList.SetHeaderColor(
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
	// 	tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	// 	tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	// 	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
	// 	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
	// 	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
	// 	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})

	// tableList.SetColumnColor(
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
	// 	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
	// )

	// tableList.SetFooterColor(
	// 	tablewriter.Colors{}, tablewriter.Colors{},
	// 	tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{},
	// 	tablewriter.Colors{tablewriter.FgHiRedColor}, tablewriter.Colors{}, tablewriter.Colors{},
	// )

	// tableList.AppendBulk(data1)
	// tableList.Render()

}
