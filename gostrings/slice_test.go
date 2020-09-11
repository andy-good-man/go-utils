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

func TestSliceUnion(t *testing.T) {
	// slice1 := []string{"1", "2", "33", "44"}
	slice2 := []string{"你好", "VIP", "咖喱鸡块和刻入啊", "44", "44", "55", "66", "44", "77"}

	// result := SliceUnion(slice1, slice2)
	result := SliceUnion(nil, slice2)

	fmt.Println("result  :", result)
}

func TestSliceIntersect(t *testing.T) {
	slice1 := []string{"1", "2", "33", "44"}
	slice2 := []string{"你好", "VIP", "咖喱鸡块和刻入啊", "44", "44", "55", "66", "44", "77"}

	result := SliceIntersect(slice1, slice2)
	fmt.Println("result  :", result)
}

func TestDifference(t *testing.T) {
	// slice1 := []string{"1", "2", "33", "44"}
	// slice2 := []string{"你好", "VIP", "咖喱鸡块和刻入啊", "44", "44", "55", "66", "44", "77"}

	// var slice1 []string
	slice2 := []string{"你好", "VIP", "咖喱鸡块和刻入啊", "44", "44", "55", "66", "44", "77"}

	// result := SliceDifference(slice1, slice2)
	// result := SliceDifference(slice1, nil)
	result := SliceDifference(nil, slice2)

	fmt.Println("result  :", result)
}

type Example struct {
	Content string
}

func (e *Example) Clone() *Example {
	res := *e
	return &res
}

func TestClone(t *testing.T) {
	r1 := new(Example)
	r1.Content = "this is example 1"
	r2 := r1.Clone()
	r2.Content = "this is example 2"

	fmt.Println(r1)
	fmt.Println(r2)

}

/*
func xxx() {
	if !found {
		err = fmt.Errorf("reverse field `%s` not found in model `%s`", fi.fullName, fi.relModelInfo.fullName)
		goto end
	}

	mForA:
		for _, ffi := range fi.relModelInfo.fields.fieldsByType[RelOneToOne] {
			if ffi.relModelInfo == mi {
				found = true
				fi.reverseField = ffi.name
				fi.reverseFieldInfo = ffi

				ffi.reverseField = fi.name
				ffi.reverseFieldInfo = fi
				break mForA
			}
		}
}

end:
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
*/
