package tmpls

// JsonReaderGo json讀取器go語言模板
var JsonReaderGo = Header + `
package {{$.Namespace}}

import (
	"encoding/json"
	"fmt"
)

// {{$.StructName}} {{$.StructNote}}
type {{$.StructName}} struct {
{{- range $.Fields}}
	{{$.FieldName .}} {{$.FieldTypeGo .}} ` + "`json:\"{{$.FieldName .}}\"`" + ` // {{$.FieldNote .}}
{{- end}}
}

// {{$.ReaderName}} {{$.StructNote}}
type {{$.ReaderName}} struct {
	Data map[{{$.PkeyGo}}]*{{$.StructName}} // 資料列表
}

// FileName 取得檔名物件
func (this *{{$.ReaderName}}) FileName() FileName {
	return NewFileName("{{$.JsonName}}", "{{$.JsonExt}}")
}

// FromData 讀取資料
func (this *{{$.ReaderName}}) FromData(data []byte) error {
	this.Data = map[{{$.PkeyGo}}]*{{$.StructName}}{}

	if err := json.Unmarshal(data, &this.Data); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	} // if

	return nil
}

// MergeData 合併資料
func (this *{{$.ReaderName}}) MergeData(data []byte) error {
	tmpl := map[{{$.PkeyGo}}]*{{$.StructName}}{}

	if err := json.Unmarshal(data, &tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	} // if

	if this.Data == nil {
		this.Data = map[{{$.PkeyGo}}]*{{$.StructName}}{}
	} // if

	for k, v := range tmpl {
		if _, ok := this.Data[k]; ok {
			return fmt.Errorf("merge data failed: key duplicate")
		} // if

		this.Data[k] = v
	} // for

	return nil
}

// Clear 清除資料
func (this *{{$.ReaderName}}) Clear() {
	this.Data = nil
}

// Get 取得資料
func (this *{{$.ReaderName}}) Get(key {{$.PkeyGo}}) *{{$.StructName}} {
	if result, ok := this.Data[key]; ok {
		return result
	} // if

	return nil
}

// Keys 取得索引列表
func (this *{{$.ReaderName}}) Keys() (result []{{$.PkeyGo}}) {
	for itor := range this.Data {
		result = append(result, itor)
	} // for

	return result
}

// Values 取得資料列表
func (this *{{$.ReaderName}}) Values() (result []*{{$.StructName}}) {
	for _, itor := range this.Data {
		result = append(result, itor)
	} // for

	return result
}

// Count 取得資料數量
func (this *{{$.ReaderName}}) Count() int {
	return len(this.Data)
}
`

// JsonSheeterGo json表格器go語言模板
var JsonSheeterGo = Header + `
package {{$.Namespace}}

// NewSheeter 建立表格資料
func NewSheeter(loader Loader) *Sheeter {
	sheeter := &Sheeter{}
	sheeter.loader = loader
	sheeter.reader = []Reader{
{{- range $.Struct}}
		&sheeter.{{.StructName}},
{{- end}}
	}
	return sheeter
}

// Sheeter 表格資料
type Sheeter struct {
	loader Loader   // 裝載器物件
	reader []Reader // 讀取器列表

{{- range $.Struct}}
	{{.StructName}} {{.ReaderName}} // {{.StructNote}}
{{- end}}
}

// FromData 讀取資料處理
func (this *Sheeter) FromData() bool {
	if this.loader == nil {
		return false
	} // if

	result := true

	for _, itor := range this.reader {
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		} // if

		if err := itor.FromData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
		} // if
	} // for

	return result
}

// MergeData 合併資料處理
func (this *Sheeter) MergeData() bool {
	if this.loader == nil {
		return false
	} // if

	result := true

	for _, itor := range this.reader {
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		} // if

		if err := itor.MergeData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
		} // if
	} // for

	return result
}

// Clear 清除資料
func (this *Sheeter) Clear() {
	for _, itor := range this.reader {
		itor.Clear()
	} // for
}

// Loader 裝載器介面
type Loader interface {
	// Load 讀取檔案
	Load(filename FileName) []byte

	// Error 錯誤處理
	Error(name string, err error)
}

// Reader 讀取器介面
type Reader interface {
	// FileName 取得檔名物件
	FileName() FileName

	// FromData 讀取資料
	FromData(data []byte) error

	// MergeData 合併資料
	MergeData(data []byte) error

	// Clear 清除資料
	Clear()
}

// NewFileName 建立檔名資料
func NewFileName(name, ext string) FileName {
	return FileName{
		name: name,
		ext:  ext,
	}
}

// FileName 檔名資料
type FileName struct {
	name string // 名稱
	ext  string // 副檔名
}

// Name 取得名稱
func (this FileName) Name() string {
	return this.name
}

// Ext 取得副檔名
func (this FileName) Ext() string {
	return this.ext
}

// File 取得完整檔名
func (this FileName) File() string {
	return this.name + this.ext
}
`
