package gostrings

import (
	"fmt"
	"testing"
)

func TestSliceIn(t *testing.T) {
	slice := []string{"1", "2", "33", "44"}
	// slice := []string{}
	// var slice []string

	fmt.Println("len(slice)  :", len(slice), slice)
	ishere := SliceIn(slice, "33")
	fmt.Println("ishere  :", ishere)
}

func TestSliceReverse(t *testing.T) {
	// slice := []string{"1", "2", "33", "44"}
	slice := []string{"你好", "VIP", "咖喱鸡块和刻入啊", "44", "44", "55", "66", "44", "77"}

	reverse := SliceReverse(slice)
	fmt.Println("reverse  :", reverse)
}
