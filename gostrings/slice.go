package gostrings

import (
	"strings"
)

// SliceIn 某值item是否在数组中
func SliceIn(slice []string, item string) bool {
	for _, val := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// SliceInLower 某值item是否在数组中, 数组元素值忽略大小写
func SliceInLower(slice []string, item string) bool {
	for _, val := range slice {
		if item == strings.ToLower(val) {
			return true
		}
	}
	return false
}

// SliceInTrimSpace 某值item是否在数组中, 数组元素值忽略首尾空格
func SliceInTrimSpace(slice []string, item string) bool {
	for _, val := range slice {
		if item == strings.TrimSpace(val) {
			return true
		}
	}
	return false
}

// SliceReverse 将数组元素顺序反向排序
func SliceReverse(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
