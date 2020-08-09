package goerrors

import "fmt"

// IfPanic 判断错误若存在，则panic
func IfPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// Errorf 格式化输出错误
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf(format, a...))
}
