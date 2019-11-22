package leetcode

// 字符串是否存在
//

func strStr(haystack string, needle string) int {
	haystackLength := len(haystack)
	needleLength := len(needle)
	if haystackLength == 0 && needleLength == 0 {
		return 0
	}
	if haystackLength == 0 || needleLength == 0 || haystackLength < needleLength {
		return -1
	}
	if haystackLength == needleLength {
		return 0
	}
	haystackArray := []rune(haystack)
	for i := 0; i < haystackLength; i++ {
		if haystackLength-i < needleLength {
			return -1
		}
		temp := string(haystackArray[i : i+needleLength])
		if temp == needle {
			return i
		}
	}
	return 0
}
