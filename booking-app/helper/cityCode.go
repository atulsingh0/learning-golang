package helper

import "strings"

func GenCityCode(city string) string {

	var cityName = strings.ToUpper(city)

	if len(city) > 3 {
		return cityName[:3]
	} else {
		padding := 3 - len(cityName)
		return cityName + strings.Repeat("z", padding)
	}
}
