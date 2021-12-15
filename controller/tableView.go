package controller

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func TableView() {

	var timess int64 = data_nCov.Results[0].UpdateTime
	// tm := time.UnixMilli(timess)
	// t := time.Now().Format("2006-01-02 15:04:05")
	tm1 := time.UnixMilli(timess).Format("2006-01-02 15:04:05")
	// fmt.Println(tm)
	// fmt.Println(tm1)
	// fmt.Println(t)

	// y := tm.Year()
	// m := tm.Month()
	// d := tm.Day()
	// h := tm.Hour()
	// i := tm.Minute()
	// s := tm.Second()
	// fmt.Println(y, m, d, h, i, s)
	// data_nCov.Results[0].Cities[localIp_city].CurrentConfirmedCount
	// var (
	// 	i                     = 1
	// 	cityName              = data_nCov.Results[0].Cities[i].CityName
	// 	currentConfirmedCount = data_nCov.Results[0].Cities[i].CurrentConfirmedCount
	// 	confirmedCount        = data_nCov.Results[0].Cities[i].ConfirmedCount
	// 	suspectedCount        = data_nCov.Results[0].Cities[i].SuspectedCount
	// 	curedCount            = data_nCov.Results[0].Cities[i].CuredCount
	// 	deadCount             = data_nCov.Results[0].Cities[i].DeadCount
	// )

	// fmt.Println("数据更新时间", "城市", "现存确诊", "累计确诊", "可疑计数", "治愈数", "死亡人数")
	// for i := 0; i < len(data_nCov.Results[0].Cities); i++ {
	// 	if data_nCov.Results[0].Cities[i].CurrentConfirmedCount != 0 {
	// 		fmt.Println(tm1,
	// 			data_nCov.Results[0].Cities[i].CityName,
	// 			data_nCov.Results[0].Cities[i].CurrentConfirmedCount,
	// 			data_nCov.Results[0].Cities[i].ConfirmedCount,
	// 			data_nCov.Results[0].Cities[i].SuspectedCount,
	// 			data_nCov.Results[0].Cities[i].CuredCount,
	// 			data_nCov.Results[0].Cities[i].DeadCount)
	// 	}
	// 	// fmt.Println(tm1, cityName, currentConfirmedCount, confirmedCount, suspectedCount, curedCount, deadCount)
	// }

	// var num1 int
	// for i := 0; i < len(data_nCov.Results[0].Cities); i++ {
	// 	num1 = i
	// }

	// 动态创建二维数组
	data2 := [][]string{}
	for i := 0; i < len(data_nCov.Results[0].Cities); i++ {
		row1 := []string{tm1, data_nCov.Results[0].Cities[i].CityName, strconv.Itoa(data_nCov.Results[0].Cities[i].CurrentConfirmedCount), strconv.Itoa(data_nCov.Results[0].Cities[i].ConfirmedCount), strconv.Itoa(data_nCov.Results[0].Cities[i].SuspectedCount), strconv.Itoa(data_nCov.Results[0].Cities[i].CuredCount), strconv.Itoa(data_nCov.Results[0].Cities[i].DeadCount)}
		data2 = append(data2, row1)
	}

	// i := 1
	// data1 := [][]string{

	// 	{tm1, data_nCov.Results[0].Cities[i].CityName, strconv.Itoa(data_nCov.Results[0].Cities[i].CurrentConfirmedCount), strconv.Itoa(data_nCov.Results[0].Cities[i].ConfirmedCount), strconv.Itoa(data_nCov.Results[0].Cities[i].SuspectedCount), strconv.Itoa(data_nCov.Results[0].Cities[i].CuredCount), strconv.Itoa(data_nCov.Results[0].Cities[i].DeadCount)},
	// 	{"12/14", "", "", "d", "d", "d", "d"},
	// 	{"12/14", "dd", "dd", "s", "h", "c", "s"},
	// }

	tableList := tablewriter.NewWriter(os.Stdout)
	tableList.SetHeader([]string{"数据更新时间", "所在地", "现存确诊", "累计确诊", "可疑计数", "治愈数", "死亡人数"})
	tableList.SetFooter([]string{tm1, data_nCov.Results[0].ProvinceName, strconv.Itoa(data_nCov.Results[0].CurrentConfirmedCount), strconv.Itoa(data_nCov.Results[0].ConfirmedCount), strconv.Itoa(data_nCov.Results[0].SuspectedCount), strconv.Itoa(data_nCov.Results[0].CuredCount), strconv.Itoa(data_nCov.Results[0].DeadCount)})
	tableList.SetBorder(false)

	tableList.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})

	tableList.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor},
	)

	tableList.SetFooterColor(
		tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiRedColor}, tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiRedColor}, tablewriter.Colors{tablewriter.FgHiRedColor}, tablewriter.Colors{tablewriter.Bold},
	)

	tableList.AppendBulk(data2)
	// tableList.AppendBulk(data1)
	tableList.Render()

}
