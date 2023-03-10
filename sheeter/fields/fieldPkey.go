package fields

import (
	"github.com/yinweli/Sheeterx/sheeter"
)

// Pkey 主要整數索引
type Pkey struct {
}

// Field 取得excel欄位類型列表
func (this *Pkey) Field() []string {
	return []string{"pkey"}
}

// IsPkey 是否是主要索引
func (this *Pkey) IsPkey() bool {
	return true
}

// ToTypeCs 取得cs類型字串
func (this *Pkey) ToTypeCs() string {
	return sheeter.TypePkeyCs
}

// ToTypeGo 取得go類型字串
func (this *Pkey) ToTypeGo() string {
	return sheeter.TypePkeyGo
}

// ToJsonValue 轉換為json值
func (this *Pkey) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil // pkey都以字串輸出, 方便json轉換
}
