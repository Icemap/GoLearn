package main

import "strings"

func SplitParam(requestUrl string) map[string]string{
	var result map[string]string
	result = make(map[string]string)
	if strings.Contains(requestUrl,"?"){
		paramString := strings.Split(requestUrl,"?")[1]
		params := strings.Split(paramString,"&")

		for i := 0 ; i < len(params);i++{
			temp := strings.Split(params[i],"=")
			result[temp[0]] = temp[1]
		}
	}

	return result
}
