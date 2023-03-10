package builds

import (
	"github.com/yinweli/Sheeterx/sheeter/excels"
)

// OperationData 作業資料
type OperationData struct {
	*Global              // 全域設定
	*excels.Excel        // excel物件
	*excels.Sheet        // sheet物件
	ExcelName     string // excel名稱
	SheetName     string // sheet名稱
}

// Operation 作業處理
func Operation(config *Config, input []*InitializeData) (result []*OperationData, err []error) {
	for _, itor := range input {
		result = append(result, &OperationData{
			Global:    &config.Global,
			Excel:     itor.Excel,
			Sheet:     itor.Sheet,
			ExcelName: itor.ExcelName,
			SheetName: itor.SheetName,
		})
	} // for

	return result, nil
}
