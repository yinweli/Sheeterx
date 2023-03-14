package nameds

import (
	"github.com/yinweli/Sheeterx/sheeter/layouts"
	"github.com/yinweli/Sheeterx/sheeter/utils"
)

// Field 欄位命名工具
type Field struct {
	Data *layouts.Data // 布局資料
}

// FieldName 取得欄位名稱
func (this *Field) FieldName() string {
	return utils.FirstUpper(this.Data.Name)
}

// FieldNote 取得欄位註解
func (this *Field) FieldNote() string {
	return this.Data.Note
}

// FieldTypeCs 取得cs欄位類型
func (this *Field) FieldTypeCs() string {
	return this.Data.Field.ToTypeCs()
}

// FieldTypeGo 取得go欄位類型
func (this *Field) FieldTypeGo() string {
	return this.Data.Field.ToTypeGo()
}
