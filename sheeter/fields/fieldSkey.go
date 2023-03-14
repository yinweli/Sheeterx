package fields

import (
	"github.com/yinweli/Sheeterx/sheeter"
)

// Skey 主要整數索引
type Skey struct {
}

// Field 取得excel欄位類型列表
func (this *Skey) Field() []string {
	return []string{"skey"}
}

// IsPkey 是否是主要索引
func (this *Skey) IsPkey() bool {
	return true
}

// ToTypeCs 取得cs類型字串
func (this *Skey) ToTypeCs() string {
	return sheeter.TypeSkeyCs
}

// ToTypeGo 取得go類型字串
func (this *Skey) ToTypeGo() string {
	return sheeter.TypeSkeyGo
}

// ToJsonValue 轉換為json值
func (this *Skey) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil // pkey都以字串輸出, 方便json轉換
}
