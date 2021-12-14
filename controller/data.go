package controller

type Data_nCov struct {
	Results []Results `json:"results"`
	Success bool      `json:"success"`
}
type Cities struct {
	CityName                 string `json:"cityName"`
	CurrentConfirmedCount    int    `json:"currentConfirmedCount"`
	ConfirmedCount           int    `json:"confirmedCount"`
	SuspectedCount           int    `json:"suspectedCount"`
	CuredCount               int    `json:"curedCount"`
	DeadCount                int    `json:"deadCount"`
	HighDangerCount          int    `json:"highDangerCount"`
	MidDangerCount           int    `json:"midDangerCount"`
	LocationID               int    `json:"locationId"`
	CurrentConfirmedCountStr string `json:"currentConfirmedCountStr"`
	CityEnglishName          string `json:"cityEnglishName,omitempty"`
}
type Results struct {
	LocationID            int         `json:"locationId"`
	ContinentName         string      `json:"continentName"`
	ContinentEnglishName  string      `json:"continentEnglishName"`
	CountryName           string      `json:"countryName"`
	CountryEnglishName    string      `json:"countryEnglishName"`
	CountryFullName       interface{} `json:"countryFullName"`
	ProvinceName          string      `json:"provinceName"`
	ProvinceEnglishName   string      `json:"provinceEnglishName"`
	ProvinceShortName     string      `json:"provinceShortName"`
	CurrentConfirmedCount int         `json:"currentConfirmedCount"` //现存确诊
	ConfirmedCount        int         `json:"confirmedCount"`        //累计确诊
	SuspectedCount        int         `json:"suspectedCount"`        //可疑计数
	CuredCount            int         `json:"curedCount"`            //治愈数
	DeadCount             int         `json:"deadCount"`             //死亡人数
	Comment               string      `json:"comment"`
	Cities                []Cities    `json:"cities"`
	UpdateTime            int64       `json:"updateTime"`
}
