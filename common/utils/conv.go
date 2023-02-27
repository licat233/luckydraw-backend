package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func HandlerIdList(ids string) []int64 {
	var idList []int64
	for _, id := range strings.Split(ids, ",") {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		idList = append(idList, idInt)
	}
	return idList
}

// expToNumber 科学计算数字转int64
// 例： ExpNum = 1.641020629e+09
func ExpToNumber(ExpNum string) (int64, error) {
	if res, err := strconv.ParseInt(ExpNum, 10, 64); err == nil {
		return res, nil
	}
	var newNum float64
	_, err := fmt.Sscanf(ExpNum, "%e", &newNum)
	if err != nil {
		return 0, fmt.Errorf("fmt.Sscanf error, ExpiresAt:%s, err:%v", ExpNum, err)
	}
	num := fmt.Sprintf("%.f", newNum)
	res, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseInt error, num:%s, err:%v", num, err)
	}
	return res, nil
}
