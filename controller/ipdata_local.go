package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Ipinfo struct {
	// IP   string `json:"ip"`
	City   string `json:"city"`
	Region string `json:"region"`
	// Country string `json:"country"`
	// Loc      string `json:"loc"`
	// Org      string `json:"org"`
	// Timezone string `json:"timezone"`
	// Readme   string `json:"readme"`
}

var dataIP Ipinfo

func Ip_local() {

	// url := "http://whois.pconline.com.cn/ipJson.jsp?json=true"
	url := "http://ipinfo.io/json"

	resp, _ := http.Get(url)
	// fmt.Println(resp)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	// fmt.Println(body)

	err := json.Unmarshal(body, &dataIP)
	if err != nil {
		fmt.Println(err)
		return
	}
	// las := gjson.Get(data, data.Pro)
	// v := simplifiedchinese.GBK.NewDecoder().Byte([]byte(data.City))
	// fmt.Println(dataIP)
	// fmt.Println(dataIP.Region)
}
