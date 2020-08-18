package gostrings

import (
	"strings"
)

// {in}.{exist}

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

// {reverse}.{sort}

// SliceReverse 将数组元素顺序反向排序
func SliceReverse(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// {union}

// SliceUnion 求并集 slice1+slice2
func SliceUnion(slice1, slice2 []string) []string {
	m := make(map[string]int) // 统计数组中某个元素的出现次数

	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}

	return slice1
}

// SliceIntersect 求交集
func SliceIntersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)

	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}

	return nn
}

// SliceDifference 求差集{A-(A和B交集)} slice1-slice2
func SliceDifference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)

	inter := SliceIntersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}

	return nn
}
