package fields

import (
	"github.com/yinweli/Sheeterx/sheeter"
)

// String 字串
type String struct {
}

// Field 取得excel欄位類型列表
func (this *String) Field() []string {
	return []string{"string"}
}

// IsPkey 是否是主要索引
func (this *String) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *String) ToTypeCs() string {
	return sheeter.TypeStringCs
}

// ToTypeGo 取得go類型字串
func (this *String) ToTypeGo() string {
	return sheeter.TypeStringGo
}

// ToJsonValue 轉換為json值
func (this *String) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil
}
