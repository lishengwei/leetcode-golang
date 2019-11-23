package leetcode

func twoSum(numbers []int, target int) []int {
	var res []int
	length := len(numbers)
	j := length - 1
	i := 0
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return append(append(res, i+1), j+1)
		} else if sum < target {
			//小于，说明有序数组的左边，已经不再满足条件了，则左边+1
			i++
		} else {
			//大于，说明有序数组右边的值大，则右边-1
			j--
		}
	}
	return res
}
