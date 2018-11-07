package utils

import "strings"

func GetUrl(url string)string  {
	if strings.HasPrefix(url,"http://") || strings.HasPrefix(url,"https://"){
		return url
	}
	return "http://"+url
}
