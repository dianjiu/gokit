package string

import (
	"strconv"
	"strings"
)

func StringToFloat64(s string) float64 {
	r, e := strconv.ParseFloat(s, 64)
	if e != nil {
		return 0
	}
	return r
}

func StrToFloat(value string) float64 {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return float
}

// 先尝试转浮点，然后再转整型
func StrToFInt(value string) int {
	float := StrToFloat(value)
	return int(float)
}

func StrToFInt64(value string) int64 {
	float := StrToFloat(value)
	return int64(float)
}

func StrToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err == nil {
		return number
	}
	return 0
}

// 取得整型值所表达的布尔类型
// 不同于golang提供的ParseBool
func StrToIntBool(value string) bool {
	// "" => false
	if len(value) <= 0 {
		return false
	}
	i := StrToFInt(value) // 先尝试尽量取得这个字符串表示的整型值
	if i <= 0 {
		return false
	}
	return true
}

// 如果是小写字母, 则变换为大写字母
func StrFirstToUpper(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// 如果是大写字母, 则变换为小写字母
func StrFirstToLower(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				return str
			} else {
				vv[i] += 32 // string的码表相差32位
				upperStr += string(vv[i])
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// 匈牙利命名法, 转换小写下划线, XxYy to xx_yy , XxYY to xx_y_y
func SnakeCasedName(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// 转换为大驼峰写法, xx_yy to XxYy
func SameCamelCaseName(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	return string(data[:])
}
