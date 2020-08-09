package gostrings

import (
	"strings"
)

// ArrayIn 某值item是否在数组中
func ArrayIn(array []string, item string) bool {
	for _, val := range array {
		if item == val {
			return true
		}
	}
	return false
}

// ArrayInLower 某值item是否在数组中, 数组元素值忽略大小写
func ArrayInLower(array []string, item string) bool {
	for _, val := range array {
		if item == strings.ToLower(val) {
			return true
		}
	}
	return false
}

// ArrayInTrimSpace 某值item是否在数组中, 数组元素值忽略首尾空格
func ArrayInTrimSpace(array []string, item string) bool {
	for _, val := range array {
		if item == strings.TrimSpace(val) {
			return true
		}
	}
	return false
}

// ArrayReverse 将数组元素顺序反向排序
func ArrayReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
