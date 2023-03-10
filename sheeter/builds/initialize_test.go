package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/sheeter/excels"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestInitialize(t *testing.T) {
	suite.Run(t, new(SuiteInitialize))
}

type SuiteInitialize struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteInitialize) SetupSuite() {
	this.Change("test-builds-initialize")
}

func (this *SuiteInitialize) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteInitialize) TestInitialize() {
	context, err := Initialize(&Config{
		Source: []string{testdata.FolderInitialize},
	})
	assert.Len(this.T(), err, 0)
	assert.Len(this.T(), context, 6)

	for _, itor := range context {
		assert.NotNil(this.T(), itor.Excel)
		assert.NotNil(this.T(), itor.Sheet)
		assert.NotEmpty(this.T(), itor.ExcelName)
		assert.NotEmpty(this.T(), itor.SheetName)
	} // for

	_, err = Initialize(&Config{
		Source: []string{testdata.FolderInitializeFailed},
	})
	assert.Len(this.T(), err, 3)
}

func (this *SuiteInitialize) TestSearchExcel() {
	result := make(chan any, sheeter.MaxExcel)
	assert.Nil(this.T(), searchExcel(testdata.FolderSearchExcel, result))
	assert.Equal(this.T(), filepath.Join(testdata.FolderSearchExcel, testdata.ExcelTest0), <-result)
	assert.Equal(this.T(), filepath.Join(testdata.FolderSearchExcel, testdata.ExcelTest1), <-result)
	assert.Equal(this.T(), filepath.Join(testdata.FolderSearchExcel, testdata.ExcelTest2), <-result)
}

func (this *SuiteInitialize) TestSearchSheet() {
	result := make(chan any, sheeter.MaxExcel)
	assert.Nil(this.T(), searchSheet(filepath.Join(testdata.FolderSearchSheet, testdata.ExcelTest0), result))
	prepare := (<-result).(*InitializeData)
	assert.NotNil(this.T(), prepare.Excel)
	assert.NotNil(this.T(), prepare.Sheet)
	assert.Equal(this.T(), testdata.ExcelTest0, prepare.ExcelName)
	assert.Equal(this.T(), testdata.SheetTest1, prepare.SheetName)
	prepare = (<-result).(*InitializeData)
	assert.NotNil(this.T(), prepare.Excel)
	assert.NotNil(this.T(), prepare.Sheet)
	assert.Equal(this.T(), testdata.ExcelTest0, prepare.ExcelName)
	assert.Equal(this.T(), testdata.SheetTest2, prepare.SheetName)
}
