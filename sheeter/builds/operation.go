package builds

import (
	"fmt"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/sheeter/excels"
	"github.com/yinweli/Sheeterx/sheeter/fields"
	"github.com/yinweli/Sheeterx/sheeter/layouts"
	"github.com/yinweli/Sheeterx/sheeter/nameds"
	"github.com/yinweli/Sheeterx/sheeter/pipelines"
	"github.com/yinweli/Sheeterx/sheeter/utils"
)

// OperationData 作業資料
type OperationData struct {
	*Global                         // 全域設定
	*excels.Excel                   // excel物件
	*excels.Sheet                   // sheet物件
	*nameds.Named                   // 命名工具
	*nameds.Pkey                    // 主要索引命名工具
	Fields          []*nameds.Field // 欄位命名工具
	*layouts.Layout                 // 布局資料
}

// Operation 作業處理
func Operation(config *Config, input []*InitializeData) (result []*OperationData, err []error) {
	for _, itor := range input {
		result = append(result, &OperationData{
			Global: &config.Global,
			Excel:  itor.Excel,
			Sheet:  itor.Sheet,
			Named: &nameds.Named{
				ExcelName: itor.ExcelName,
				SheetName: itor.SheetName,
			},
		})
	} // for

	_, err = pipelines.Pipeline[*OperationData]("search excel", result, []pipelines.PipelineFunc[*OperationData]{
		parseLayout,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	return result, nil
}

// parseLayout 解析布局
func parseLayout(input *OperationData, _ chan any) error {
	line, err := input.GetLine(
		input.SheetName,
		input.LineOfTag,
		input.LineOfName,
		input.LineOfNote,
		input.LineOfField,
	)

	if err != nil {
		return fmt.Errorf("parse layout: %w: %v#%v", err, input.ExcelName, input.SheetName)
	} // if

	layout := layouts.NewLayout()
	lineTag := line[input.LineOfTag][sheeter.OutputCol:]   // 尋訪時, 以標籤行為主
	lineName := line[input.LineOfName][sheeter.OutputCol:] // 第sheeter.OutputCol欄之後才是資料欄, 因此在前面過濾掉
	lineNote := line[input.LineOfNote][sheeter.OutputCol:]
	lineField := line[input.LineOfField][sheeter.OutputCol:]

	for col, itor := range lineTag {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		tag := itor
		name := utils.GetItem(lineName, col)
		note := utils.GetItem(lineNote, col)
		field, err := fields.Parser(utils.GetItem(lineField, col))

		if err != nil {
			return fmt.Errorf("parse layout: %w: %v#%v(column:%v)", err, input.ExcelName, input.SheetName, col)
		} // if

		if err := layout.Add(tag, name, note, field); err != nil {
			return fmt.Errorf("parse layout: %w: %v#%v(column:%v)", err, input.ExcelName, input.SheetName, col)
		} // if
	} // for

	pkey := layout.Pkey(input.Tag)

	if pkey == nil {
		return fmt.Errorf("parse layout: pkey not exist: %v#%v", input.ExcelName, input.SheetName)
	} // if

	input.Pkey = &nameds.Pkey{
		Pkey: pkey.Field,
	}

	for _, itor := range layout.Layout(input.Tag) {
		input.Fields = append(input.Fields, &nameds.Field{
			Data: itor,
		})
	} // for

	input.Layout = layout
	return nil
}
