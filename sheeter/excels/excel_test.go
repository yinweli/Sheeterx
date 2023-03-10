package excels

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/testdata"
)

func TestExcel(t *testing.T) {
	suite.Run(t, new(SuiteExcel))
}

type SuiteExcel struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteExcel) SetupSuite() {
	this.Change("test-excels-excel")
}

func (this *SuiteExcel) TearDownSuite() {
	CloseAll()
	this.Restore()
}

func (this *SuiteExcel) TestOpen() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelTest0))
	assert.True(this.T(), target.IsOpen())
	target.Close()
	assert.False(this.T(), target.IsOpen())

	target = this.target()
	assert.NotNil(this.T(), target.Open(testdata.UnknownStr))

	CloseAll()
}

func (this *SuiteExcel) TestGet() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelTest0))

	sheet, err := target.Get(testdata.SheetTest1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	sheet.Close()

	sheet, err = target.Get(testdata.SheetTest2)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), sheet)
	sheet.Close()

	_, err = target.Get(testdata.UnknownStr)
	assert.NotNil(this.T(), err)

	CloseAll()
}

func (this *SuiteExcel) TestGetLine() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelTest0))

	line, err := target.GetLine(testdata.SheetTest1, 1, 2, 3)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	assert.Len(this.T(), line, 3)
	assert.Equal(this.T(), []string{"value1", "value2"}, line[1])
	assert.Equal(this.T(), []string{"value3", "value4"}, line[2])
	assert.Equal(this.T(), []string{}, line[3])

	line, err = target.GetLine(testdata.SheetTest2, 1, 2, 3)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	assert.Len(this.T(), line, 3)
	assert.Equal(this.T(), []string{"value5", "value6"}, line[1])
	assert.Equal(this.T(), []string{"value7", "value8"}, line[2])
	assert.Equal(this.T(), []string{}, line[3])

	_, err = target.GetLine(testdata.SheetTest1, -1)
	assert.NotNil(this.T(), err)

	_, err = target.GetLine(testdata.UnknownStr, 1)
	assert.NotNil(this.T(), err)

	CloseAll()
}

func (this *SuiteExcel) TestSheets() {
	target := this.target()
	assert.Equal(this.T(), []string{}, target.Sheets())
	assert.Nil(this.T(), target.Open(testdata.ExcelTest0))
	assert.Equal(this.T(), []string{testdata.SheetTest1, testdata.SheetTest2}, target.Sheets())
	CloseAll()
}

func (this *SuiteExcel) TestExist() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelTest0))
	assert.True(this.T(), target.Exist(testdata.SheetTest1))
	assert.True(this.T(), target.Exist(testdata.SheetTest2))
	assert.False(this.T(), target.Exist(testdata.UnknownStr))
	CloseAll()
}

func (this *SuiteExcel) TestSheet() {
	target := this.target()
	assert.Nil(this.T(), target.Open(testdata.ExcelTest0))

	sheet, err := target.Get(testdata.SheetTest1)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Next())
	data, err := sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value1", "value2"}, data)
	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value3", "value4"}, data)

	sheet, err = target.Get(testdata.SheetTest2)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value5", "value6"}, data)
	assert.True(this.T(), sheet.Next())
	data, err = sheet.Data()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"value7", "value8"}, data)

	sheet, err = target.Get(testdata.SheetTest1)
	assert.Nil(this.T(), err)
	assert.True(this.T(), sheet.Nextn(2))
	assert.False(this.T(), sheet.Nextn(-1))
	sheet.Close()

	sheet, err = target.Get(testdata.SheetTest1)
	assert.Nil(this.T(), err)
	_, err = sheet.Data()
	assert.NotNil(this.T(), err)
	sheet.Close()

	CloseAll()
}

func (this *SuiteExcel) TestColumnToIndex() {
	assert.Equal(this.T(), 1, columnToIndex("A"))
	assert.Equal(this.T(), 111, columnToIndex("DG"))
	assert.Equal(this.T(), 222, columnToIndex("HN"))
	assert.Equal(this.T(), 333, columnToIndex("LU"))
	assert.Equal(this.T(), 444, columnToIndex("QB"))
	assert.Equal(this.T(), 555, columnToIndex("UI"))
	assert.Equal(this.T(), 666, columnToIndex("YP"))
	assert.Equal(this.T(), 777, columnToIndex("ACW"))
	assert.Panics(this.T(), func() { columnToIndex("") })
	assert.Panics(this.T(), func() { columnToIndex("0") })
	assert.Panics(this.T(), func() { columnToIndex("?") })
}

func (this *SuiteExcel) target() *Excel {
	return &Excel{}
}
