package leetcode

//二进制字符串相加

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var sumArray []int
	aArray := []rune(a)
	bArray := []rune(b)
	aLength := len(aArray)
	bLength := len(bArray)
	temp := 0
	i := 1
	for true {
		aKey := aLength - i
		bKey := bLength - i
		if aKey < 0 && bKey < 0 {
			break
		}
		var aItem = 0
		var bItem = 0
		if aKey >= 0 {
			aItem, _ = strconv.Atoi(string(aArray[aKey]))
		}
		if bKey >= 0 {
			bItem, _ = strconv.Atoi(string(bArray[bKey]))
		}
		sum := aItem + bItem + temp
		if sum > 1 {
			temp = 1
			sumArray = append(sumArray, sum % 2)
		} else {
			temp = 0
			sumArray = append(sumArray, sum)
		}
		i++
	}
	if temp > 0 {
		sumArray = append(sumArray, 1)
	}
	sumString := ""
	for i := len(sumArray) - 1; i >= 0; i-- {
		sumString += strconv.Itoa(sumArray[i])
	}
	return sumString
}
