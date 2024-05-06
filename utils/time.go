package utils

import (
	"strconv"
	"strings"
	"time"
)

// 支持的格式
// 1. 7d33m
// 2. 1234567
// 3. 33m
func ParseExpiresTime(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		day, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(day)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
