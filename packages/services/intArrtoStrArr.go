package services

import "strconv"

func IntArrToStrArr(intArr []int) []string {

	strArr := []string{}
	for _, item := range intArr {
		strArr = append(strArr, strconv.Itoa(item))
	}

	return strArr
}
