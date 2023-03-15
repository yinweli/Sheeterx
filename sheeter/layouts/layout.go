package layouts

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/sheeter/fields"
	"github.com/yinweli/Sheeterx/sheeter/utils"
)

// NewLayout 建立資料布局器
func NewLayout() *Layout {
	return &Layout{
		layout: map[int]*Data{},
	}
}

// Layout 資料布局器
type Layout struct {
	layout map[int]*Data // 布局列表
	pkey   *Data         // 主要索引資料
}

// Data 布局資料
type Data struct {
	Tag   string       // 標籤字串
	Name  string       // 欄位名稱
	Note  string       // 欄位註解
	Field fields.Field // 欄位類型
}

// Set 設定布局
func (this *Layout) Set(lineTag, lineName, lineNote, lineField []string) error {
	result := []string{}

	for col, itor := range lineTag { // 尋訪時, 以標籤行為主
		if col == sheeter.OutputIndex { // 跳過輸出欄
			continue
		} // if

		if itor == "" { // 一旦遇到空欄位, 就結束布局
			break
		} // if

		name := utils.GetItem(lineName, col)
		note := utils.GetItem(lineNote, col)
		field := utils.GetItem(lineField, col)

		if err := this.add(col, itor, name, note, field); err != nil {
			result = append(result, fmt.Sprintf("[column %v: %v]", col, err.Error()))
		} // if
	} // for

	if len(result) > 0 {
		return fmt.Errorf("layout set: %v", strings.Join(result, ", "))
	} // if

	return nil
}

// Pack 打包資料
func (this *Layout) Pack(tag string, data []string) (result map[string]interface{}, pkey any, err error) {
	if output := utils.GetItem(data, sheeter.OutputIndex); utils.CheckIgnore(output) { // 輸出欄檢查
		return nil, nil, nil
	} // if

	result = map[string]interface{}{}

	for col, itor := range this.layout {
		if utils.CheckTag(tag, itor.Tag) == false {
			continue
		} // if

		value, err := itor.Field.ToJsonValue(utils.GetItem(data, col))

		if err != nil {
			return nil, nil, fmt.Errorf("layout pack: %w", err)
		} // if

		if itor.Field.IsPkey() {
			pkey = value
		} // if

		result[itor.Name] = value
	} // for

	return result, pkey, nil
}

// Layout 取得布局資料
func (this *Layout) Layout(tag string) (result []*Data) {
	for _, itor := range this.layout {
		if utils.CheckTag(tag, itor.Tag) {
			result = append(result, itor)
		} // if
	} // for

	return result
}

// Pkey 取得主要索引資料
func (this *Layout) Pkey(tag string) *Data {
	if this.pkey != nil && utils.CheckTag(tag, this.pkey.Tag) {
		return this.pkey
	} // if

	return nil
}

// add 新增布局
func (this *Layout) add(col int, tag, name, note, field string) error {
	if name == "" {
		return fmt.Errorf("layout add: name empty")
	} // if

	if utils.CheckName(name) == false {
		return fmt.Errorf("layout add: name invalid")
	} // if

	for _, itor := range this.layout {
		if itor.Name == name {
			return fmt.Errorf("layout add: name duplicate")
		} // if
	} // if

	field_, err := fields.Parser(field)

	if err != nil {
		return fmt.Errorf("layout add: %w", err)
	} // if

	data := &Data{
		Tag:   tag,
		Name:  name,
		Note:  note,
		Field: field_,
	}

	if field_.IsPkey() {
		if this.pkey != nil {
			return fmt.Errorf("layout add: too many pkey")
		} // if

		this.pkey = data
	} // if

	this.layout[col] = data
	return nil
}
