package leetcode

import (
	"reflect"
	"testing"
)

//两数之和 II - 输入有序数组
//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
//说明:
//
//返回的下标值（index1 和 index2）不是从零开始的。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
func TestTowSum(t *testing.T) {
	numbers := []int{2,7,11,15}
	target := 9
	got := twoSum(numbers, target)
	want := []int{1,2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("error")
	}
}