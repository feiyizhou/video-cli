package utils

import (
	"strconv"
	"strings"
)

func ZeroStrToInt(str string) (int, error) {
	if strings.EqualFold(str, "00") {
		return 0, nil
	}
	if strings.HasPrefix(str, "0") {
		return ZeroStrToInt(str[1:])
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, err
}

func TimeStrToSeconds(timeStr string) (int, error) {
	timeStrArr := strings.Split(timeStr, ":")
	hours, err := ZeroStrToInt(timeStrArr[0])
	if err != nil {
		return 0, err
	}
	minutes, err := ZeroStrToInt(timeStrArr[1])
	if err != nil {
		return 0, err
	}
	seconds, err := ZeroStrToInt(timeStrArr[2])
	if err != nil {
		return 0, err
	}
	seconds += hours * 3600
	seconds += minutes * 60
	return seconds, err
}
