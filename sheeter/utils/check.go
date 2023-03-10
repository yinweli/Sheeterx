package utils

import (
	"strings"

	"github.com/yinweli/Sheeterx/sheeter"
)

// CheckKeyword 關鍵字檢查
func CheckKeyword(input string) bool {
	for _, itor := range sheeter.Keyword {
		if strings.EqualFold(input, itor) {
			return false
		} // if
	} // for

	return true
}

// CheckIgnore 忽略檢查
func CheckIgnore(input string) bool {
	return strings.HasPrefix(strings.ToLower(input), strings.ToLower(sheeter.TokenIgnore))
}

// CheckName 名稱檢查
func CheckName(input string) bool {
	if input == "" { // 名稱不能為空
		return false
	} // if

	if input[0] >= '0' && input[0] <= '9' { // 名稱不能以數字開頭
		return false
	} // if

	for _, itor := range input { // 名稱必須是字母, 數字與'_'的組合
		if (itor < 'a' || itor > 'z') && (itor < 'A' || itor > 'Z') && (itor < '0' || itor > '9') && itor != '_' {
			return false
		} // if
	} // for

	return true
}

// CheckTag 標籤檢查
func CheckTag(input, tag string) bool {
	if CheckIgnore(tag) {
		return false
	} // if

	for _, itor := range input {
		if strings.ContainsRune(tag, itor) {
			return true
		} // if
	} // for

	return false
}
