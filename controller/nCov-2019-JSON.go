package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var data_nCov Data_nCov

func PrintJSON() {
	// https://lab.isaaclin.cn/nCoV/api/provinceName?lang=en
	// https://lab.isaaclin.cn/nCoV/api/area?provinceEng=Zhejiang&latest=true
	// https://lab.isaaclin.cn/nCoV/api/area?province=%E6%B5%99%E6%B1%9F%E7%9C%81&latest=true
	params := url.Values{}
	Url, err := url.Parse("https://lab.isaaclin.cn/nCoV/api/area")
	// defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	params.Set("provinceEng", dataIP.Region)
	params.Set("latest", "true")

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	// fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// if resp.StatusCode == 200 {
	// 	fmt.Println("Status OK , 200")
	// }

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &data_nCov)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(data_nCov)
	// fmt.Println(data.Results.)

	// fmt.Println(data_nCov.Results[0].ProvinceName)              //浙江省
	// fmt.Println(data_nCov.Results[0].Cities[0].CityName)        //绍兴
	// fmt.Println(data_nCov.Results[0].Cities[0].CityEnglishName) //Shaoxing
	// fmt.Println(dataIP.City)                                    //Ningbo
	// fmt.Println(len(data_nCov.Results[0].Cities))

}
