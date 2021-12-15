package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NewData struct {
	ResultsNew []ResultsNew `json:"results"`
	// Success    bool         `json:"success"`
}

type ResultsNew struct {
	// PubDate    string      `json:"pubDate"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	InfoSource string `json:"infoSource"`
	// SourceURL  string      `json:"sourceUrl"`
	// Province   interface{} `json:"province"`
	// ProvinceID string      `json:"provinceId"`
}

var newData NewData

func NewView() {
	url := "https://lab.isaaclin.cn/nCoV/api/news?page=1&num=10"
	resp2, _ := http.Get(url)
	body2, _ := ioutil.ReadAll(resp2.Body)
	defer resp2.Body.Close()
	err1 := json.Unmarshal(body2, &newData)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	// fmt.Println(newData)

	// for i := 0; i < len(newData.ResultsNew[0].Title); i++ {
	fmt.Println("最近的3条疫情相关新闻: 👇👇👇")
	for i := 0; i < 3; i++ {
		fmt.Printf("%c[33;47;7m%s%c[0m", 0x1B, "新闻标题: ", 0x1B)
		fmt.Println(newData.ResultsNew[i].Title)
		fmt.Printf("%c[34;47;7m%s%c[0m", 0x1B, "新闻内容: ", 0x1B)
		fmt.Println(newData.ResultsNew[i].Summary)
		fmt.Printf("%c[35;47;7m%s%c[0m", 0x1B, "新闻来源: ", 0x1B)
		fmt.Println(newData.ResultsNew[i].InfoSource)
		fmt.Println("")
		fmt.Println("✨✨✨✨✨---------------------正义分割线---------------------✨✨✨✨✨")
		fmt.Println("")
	}
}
