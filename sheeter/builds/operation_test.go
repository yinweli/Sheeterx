package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter/excels"
	"github.com/yinweli/Sheeterx/sheeter/fields"
	"github.com/yinweli/Sheeterx/sheeter/nameds"
	"github.com/yinweli/Sheeterx/sheeter/utils"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestOperation(t *testing.T) {
	suite.Run(t, new(SuiteOperation))
}

type SuiteOperation struct {
	suite.Suite
	testdata.TestData
	excel              string
	sheetSuccess       string
	sheetFieldNotExist string
	sheetFieldEmpty    string
	sheetPkeyNotExist  string
	lineOfTag          int
	lineOfName         int
	lineOfNote         int
	lineOfField        int
	lintOfData         int
}

func (this *SuiteOperation) SetupSuite() {
	this.TBegin("test-builds-operation", "operation")
	this.excel = "excel.xlsx"
	this.sheetSuccess = "Success"
	this.sheetFieldNotExist = "Failed1"
	this.sheetFieldEmpty = "Failed2"
	this.sheetPkeyNotExist = "Failed3"
	this.lineOfTag = 1
	this.lineOfName = 2
	this.lineOfNote = 3
	this.lineOfField = 4
	this.lintOfData = 5
}

func (this *SuiteOperation) TearDownSuite() {
	excels.CloseAll()
	this.TFinal()
}

func (this *SuiteOperation) TestParseLayout() {
	operationData := this.prepare(this.excel, this.sheetSuccess, "1")
	assert.Nil(this.T(), parseLayout(operationData, nil))
	assert.NotNil(this.T(), operationData.Pkey)
	assert.Equal(this.T(), &fields.Pkey{}, operationData.Pkey.Pkey)
	assert.NotNil(this.T(), operationData.Fields)
	assert.Len(this.T(), operationData.Fields, 5)
	assert.NotNil(this.T(), operationData.Layout)

	operationData.SheetName = this.Unknown
	assert.NotNil(this.T(), parseLayout(operationData, nil))

	operationData.SheetName = this.sheetFieldNotExist
	assert.NotNil(this.T(), parseLayout(operationData, nil))

	operationData.SheetName = this.sheetFieldEmpty
	assert.NotNil(this.T(), parseLayout(operationData, nil))

	operationData.SheetName = this.sheetPkeyNotExist
	assert.NotNil(this.T(), parseLayout(operationData, nil))
}

func (this *SuiteOperation) TestGenerateData() {
	expected, _ := utils.JsonMarshal(map[string]interface{}{
		"1": map[string]interface{}{
			"pkey":  int32(1),
			"name1": int32(10),
			"name2": int32(11),
			"name3": int32(12),
			"name4": int32(13),
		},
		"2": map[string]interface{}{
			"pkey":  int32(2),
			"name1": int32(20),
			"name2": int32(21),
			"name3": int32(22),
			"name4": int32(23),
		},
		"4": map[string]interface{}{
			"pkey":  int32(4),
			"name1": int32(40),
			"name2": int32(41),
			"name3": int32(42),
			"name4": int32(43),
		},
		"5": map[string]interface{}{
			"pkey":  int32(5),
			"name1": int32(50),
			"name2": int32(51),
			"name3": int32(52),
			"name4": int32(53),
		},
	})

	operationData := this.prepare(this.excel, this.sheetSuccess, "1")
	assert.Nil(this.T(), parseLayout(operationData, nil))
	assert.Nil(this.T(), generateData(operationData, nil))
	this.AssertCompareFile(this.T(), operationData.DataPath(), expected)
}

func (this *SuiteOperation) prepare(excelName, sheetName, tag string) *OperationData {
	excel := &excels.Excel{}
	assert.Nil(this.T(), excel.Open(excelName))
	sheet, err := excel.Get(sheetName)
	assert.Nil(this.T(), err)
	return &OperationData{
		Global: &Global{
			LineOfTag:   this.lineOfTag,
			LineOfName:  this.lineOfName,
			LineOfNote:  this.lineOfNote,
			LineOfField: this.lineOfField,
			LineOfData:  this.lintOfData,
			Tag:         tag,
		},
		Excel: excel,
		Sheet: sheet,
		Named: &nameds.Named{
			ExcelName: excelName,
			SheetName: sheetName,
		},
	}
}
