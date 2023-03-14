package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter/excels"
	"github.com/yinweli/Sheeterx/sheeter/fields"
	"github.com/yinweli/Sheeterx/sheeter/nameds"
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
}

func (this *SuiteOperation) SetupSuite() {
	this.TBegin("test-builds-operation", "operation")
	this.excel = "excel.xlsx"
	this.sheetSuccess = "Success"
	this.sheetFieldNotExist = "Failed1"
	this.sheetFieldEmpty = "Failed2"
	this.sheetPkeyNotExist = "Failed3"
}

func (this *SuiteOperation) TearDownSuite() {
	excels.CloseAll()
	this.TFinal()
}

func (this *SuiteOperation) TestParseLayout() {
	excel := &excels.Excel{}
	assert.Nil(this.T(), excel.Open(this.excel))
	operationData := &OperationData{
		Global: &Global{
			LineOfTag:   1,
			LineOfName:  2,
			LineOfNote:  3,
			LineOfField: 4,
			LineOfData:  5,
			Tag:         "OK",
		},
		Excel: excel,
		Sheet: nil,
		Named: &nameds.Named{
			ExcelName: this.excel,
			SheetName: this.sheetSuccess,
		},
	}

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
