package gostrings

// ArrayInExist 某值item是否在数组中
func ArrayInExist(array []string, item string) bool {
	for _, val := range array {
		if val == item {
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
