package convert

import "strconv"

// UtoA uint 转 string
func UtoA(val uint) string {
	return strconv.Itoa(int(val))
}

// Join 用 split 作为分割符 组合str
func Join(split string, str ...string) string {
	if len(str) == 0 {
		return ""
	}
	res := str[0]
	for i := 1; i < len(str); i++ {
		res += ":" + str[i]
	}
	return res
}
