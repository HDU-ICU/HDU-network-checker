package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type ApiResult struct {
	Error             string `json:"error"`
	UserName          string `json:"user_name"`
	ProductsId        string `json:"products_id"`
	ProductsName      string `json:"products_name"`
	OnlineDeviceTotal string `json:"online_device_total"`
}

func GetAAAInfo(baseUrl string) (*ApiResult, error) {
	api := baseUrl + "/cgi-bin/rad_user_info?callback=jQuery"
	resp, err := Get(api)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`jQuery\((.*)\)`)

	if re.MatchString(resp) {
		match := re.FindStringSubmatch(resp)
		if len(match) == 2 {
			var result ApiResult
			err := json.Unmarshal([]byte(match[1]), &result)
			if err != nil {
				return nil, err
			}
			return &result, nil
		}
	}

	return nil, fmt.Errorf("API 返回格式错误+%v", resp)
}
