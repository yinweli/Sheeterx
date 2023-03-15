package fields

import (
	"fmt"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/sheeter/utils"
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
	result, err = utils.StrToInt32(input)

	if err != nil {
		return nil, fmt.Errorf("pkey to json value: %w", err)
	} // if

	return result, nil
}
