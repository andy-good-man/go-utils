package gostrings

import (
	"fmt"
	"testing"
)

func TestArrayIn(t *testing.T) {
	array := []string{"1", "2", "33", "44"}
	// array := []string{}
	// var array []string

	fmt.Println("len(array)  :", len(array), array)
	ishere := ArrayIn(array, "33")
	fmt.Println("ishere  :", ishere)
}

func TestReverse(t *testing.T) {
	// array := []string{"1", "2", "33", "44"}
	array := []string{"你好", "VIP", "咖喱鸡块和刻入啊", "44", "44", "55", "66", "44", "77"}

	reverse := ArrayReverse(array)
	fmt.Println("reverse  :", reverse)
}
