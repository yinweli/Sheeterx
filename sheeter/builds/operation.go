package builds

import (
	"fmt"

	"github.com/yinweli/Sheeterx/sheeter/excels"
	"github.com/yinweli/Sheeterx/sheeter/layouts"
	"github.com/yinweli/Sheeterx/sheeter/nameds"
	"github.com/yinweli/Sheeterx/sheeter/pipelines"
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
		return fmt.Errorf("parse layout: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	layout := layouts.NewLayout()
	lineTag := line[input.LineOfTag]
	lineName := line[input.LineOfName]
	lineNote := line[input.LineOfNote]
	lineField := line[input.LineOfField]

	if err = layout.Set(lineTag, lineName, lineNote, lineField); err != nil {
		return fmt.Errorf("parse layout: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	pkey := layout.Pkey(input.Tag)

	if pkey == nil {
		return fmt.Errorf("parse layout: %v#%v: pkey not exist", input.ExcelName, input.SheetName)
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
