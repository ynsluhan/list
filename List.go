package List

import (
	"regexp"
)

/**
* @Author: yNsLuHan
* @Description:
* @File: ListIn
* @Version: 1.0.0
* @Date: 2021/6/9 4:20 下午
 */
func StringInList(PatternList []string, s string) bool {
	for _, url := range PatternList {
		if url == s {
			return true
		}
	}
	return false
}

/**
 * @Author yNsLuHan
 * @Description: 判断int类型是否在列表中
 * @Time 2021-06-10 16:22:19
 * @param PatternList
 * @param id
 * @return bool
 */
func IntInList(PatternList []int, id int) bool {
	for _, patternId := range PatternList {
		if patternId == id {
			return true
		}
	}
	return false
}

/**
 * @Author yNsLuHan
 * @Description: 判断url是否在urlList中，正则匹配
 * @Time 2021-06-09 16:22:43
 * @param PatternList
 * @param url
 * @return bool
 */
func UrlInList(PatternList []string, s string) bool {
	for _, url := range PatternList {
		matchString, _ := regexp.MatchString(url, s)
		if matchString {
			return true
		}
	}
	return false
}