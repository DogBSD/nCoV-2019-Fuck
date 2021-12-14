package controller

import "fmt"

func MathingCity() int {
	var localipCity int
	for i := 0; i < len(data_nCov.Results[0].Cities); i++ {
		if data_nCov.Results[0].Cities[i].CityEnglishName == dataIP.City {
			fmt.Println(data_nCov.Results[0].Cities[i].CityName)
			localipCity = i
		}
	}
	return localipCity
}
