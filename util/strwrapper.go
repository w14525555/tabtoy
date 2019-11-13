package util

import (
	"fmt"
	"strings"
)

func StringEscape(s string) string {

	b := make([]byte, 0)

	var index int

	// 表中直接使用换行会干扰最终合并文件格式, 所以转成\n,由pbt文本解析层转回去
	for index < len(s) {
		c := s[index]

		switch c {
		// case '"':
		// 	b = append(b, '\\')
		// 	b = append(b, '"')
		case '\n':
			b = append(b, '\\')
			b = append(b, 'n')
		case '\r':
			b = append(b, '\\')
			b = append(b, 'r')
		// case '{':
		// 	fmt.Println("111111111111")
		// 	b = append(b, '{')
		// case '\\':

		// 	var nextChar byte
		// 	if index+1 < len(s) {
		// 		nextChar = s[index+1]
		// 	}

		// 	b = append(b, '\\')

		// 	switch nextChar {
		// 	case 'n', 'r':
		// 	default:
		// 		b = append(b, c)
		// 	}
		default:
			b = append(b, c)
		}

		index++

	}

	// 如果是{}类型的节点 说明是二维数组自己加的节点 不能返回引号形式
	var strB = string(b)
	if strB == "{" || strB == "}" {
		return strB
	}

	// 双引号返回
	var length = len(strB)
	if length > 0 && strB[0] == '"' && strB[length-1] == '"' {
		return strB
	}

	// 单引号也直接返回
	if length > 0 && rune(strB[0]) == 39 {
		return strB
	}

	// 数字直接返回数字
	if length > 0 && isDigits(strB) {
		return strB
	}

	return fmt.Sprintf("'%s'", strB)

}

func isDigits(str string) bool {
	if str == "" {
		return false
	}

	for _, v := range str {
		if !strings.Contains("0123456789", string(v)) {
			return false
		}
	}
	return true
}

// 单值的字符串处理
func StringEscapeForSingle(s string) string {

	b := make([]byte, 0)

	var index int

	// 表中直接使用换行会干扰最终合并文件格式, 所以转成\n,由pbt文本解析层转回去
	for index < len(s) {
		c := s[index]

		switch c {
		// case '"':
		// 	b = append(b, '\\')
		// 	b = append(b, '"')
		case '\n':
			b = append(b, '\\')
			b = append(b, 'n')
		case '\r':
			b = append(b, '\\')
			b = append(b, 'r')
		// case '{':
		// 	fmt.Println("111111111111")
		// 	b = append(b, '{')
		// case '\\':

		// 	var nextChar byte
		// 	if index+1 < len(s) {
		// 		nextChar = s[index+1]
		// 	}

		// 	b = append(b, '\\')

		// 	switch nextChar {
		// 	case 'n', 'r':
		// 	default:
		// 		b = append(b, c)
		// 	}
		default:
			b = append(b, c)
		}

		index++

	}

	// 如果是{}类型的节点 说明是二维数组自己加的节点 不能返回引号形式
	var strB = string(b)
	if strB == "{" || strB == "}" {
		return strB
	}

	// // 双引号返回
	var length = len(strB)
	// if length > 0 && strB[0] == '"' && strB[length-1] == '"' {
	// 	return strB
	// }

	// 单引号也直接返回
	if length > 0 && rune(strB[0]) == 39 {
		return strB
	}

	return fmt.Sprintf("'%s'", strB)

}
