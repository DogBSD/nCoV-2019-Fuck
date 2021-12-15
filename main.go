package main

import "gitee.com/CatBSD/nCoV-2019-Fuck/controller"

func main() {
	// fmt.Println("COVID-2019")
	// https://lab.isaaclin.cn/nCoV/api/provinceName?lang=en
	// https://lab.isaaclin.cn/nCoV/api/area?provinceEng=Zhejiang&latest=true
	// https://lab.isaaclin.cn/nCoV/api/area?province=%E6%B5%99%E6%B1%9F%E7%9C%81&latest=true
	//打印程序的Logo，同时请求IP和相应的城市疫情数据，完成后晴空屏幕
	controller.PrintLogo()
	// 获取用户IP
	controller.Ip_local()

	//所在地城市的疫情状况
	controller.PrintJSON()

	//将疫情数据呈现表格
	controller.TableView()

	// 最近的新闻
	controller.NewView()

	// fmt.Println(controller.MathingCity())
}
