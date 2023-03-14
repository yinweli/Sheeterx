package layouts

import (
	"fmt"

	"github.com/yinweli/Sheeterx/sheeter/fields"
	"github.com/yinweli/Sheeterx/sheeter/utils"
)

// NewLayout 建立資料布局器
func NewLayout() *Layout {
	return &Layout{
		layout: []*Data{},
	}
}

// Layout 資料布局器
type Layout struct {
	layout []*Data // 布局列表
	pkey   *Data   // 主要索引資料
}

// Data 布局資料
type Data struct {
	Tag   string       // 標籤字串
	Name  string       // 欄位名稱
	Note  string       // 欄位註解
	Field fields.Field // 欄位類型
}

// Add 新增布局
func (this *Layout) Add(tag, name, note string, field fields.Field) error {
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

	if field == nil {
		return fmt.Errorf("layout add: field nil")
	} // if

	if this.pkey != nil && field.IsPkey() {
		return fmt.Errorf("layout add: too many pkey")
	} // if

	data := &Data{
		Tag:   tag,
		Name:  name,
		Note:  note,
		Field: field,
	}
	this.layout = append(this.layout, data)

	if field.IsPkey() {
		this.pkey = data
	} // if

	return nil
}

// Pack 打包資料
func (this *Layout) Pack(tag string, input []string) (result map[string]interface{}, pkey any, err error) {
	result = map[string]interface{}{}

	for i, itor := range this.layout {
		if utils.CheckTag(tag, itor.Tag) == false {
			continue
		} // if

		value, err := itor.Field.ToJsonValue(utils.GetItem(input, i))

		if err != nil {
			return nil, 0, fmt.Errorf("layout pack: %w", err)
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
