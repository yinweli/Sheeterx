package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter/excels"
	"github.com/yinweli/Sheeterx/sheeter/utils"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestJsonPack(t *testing.T) {
	suite.Run(t, new(SuiteJsonPack))
}

type SuiteJsonPack struct {
	suite.Suite
	testdata.TestData
	excel        string
	sheetSuccess string
	sheetFailed  string
	sheetEmpty   string
	lineOfTag    int
	lineOfName   int
	lineOfNote   int
	lineOfField  int
	lintOfData   int
}

func (this *SuiteJsonPack) SetupSuite() {
	this.TBegin("test-layouts-jsonPack", "jsonPack")
	this.excel = "excel.xlsx"
	this.sheetSuccess = "Success"
	this.sheetFailed = "Failed"
	this.sheetEmpty = "Empty"
	this.lineOfTag = 1
	this.lineOfName = 2
	this.lineOfNote = 3
	this.lineOfField = 4
	this.lintOfData = 5
}

func (this *SuiteJsonPack) TearDownSuite() {
	excels.CloseAll()
	this.TFinal()
}

func (this *SuiteJsonPack) TestJsonPack() {
	expected1, _ := utils.JsonMarshal(map[string]interface{}{
		"1": map[string]interface{}{
			"pkey":  int32(1),
			"name1": int32(1),
			"name2": []int32{10, 11, 12},
			"name3": "a",
			"name4": []string{"a", "b", "c", "d"},
		},
		"2": map[string]interface{}{
			"pkey":  int32(2),
			"name1": int32(2),
			"name2": []int32{20, 21, 22},
			"name3": "b",
			"name4": []string{"b", "c", "d", "a"},
		},
		"3": map[string]interface{}{
			"pkey":  int32(3),
			"name1": int32(3),
			"name2": []int32{30, 31, 32},
			"name3": "c",
			"name4": []string{"c", "d", "a", "b"},
		},
	})
	expected2, _ := utils.JsonMarshal(map[string]interface{}{
		"1": map[string]interface{}{
			"pkey":  int32(1),
			"name1": int32(1),
			"name2": []int32{10, 11, 12},
		},
		"2": map[string]interface{}{
			"pkey":  int32(2),
			"name1": int32(2),
			"name2": []int32{20, 21, 22},
		},
		"3": map[string]interface{}{
			"pkey":  int32(3),
			"name1": int32(3),
			"name2": []int32{30, 31, 32},
		},
	})
	expected3, _ := utils.JsonMarshal(map[string]interface{}{
		"1": map[string]interface{}{
			"pkey":  int32(1),
			"name3": "a",
			"name4": []string{"a", "b", "c", "d"},
		},
		"2": map[string]interface{}{
			"pkey":  int32(2),
			"name3": "b",
			"name4": []string{"b", "c", "d", "a"},
		},
		"3": map[string]interface{}{
			"pkey":  int32(3),
			"name3": "c",
			"name4": []string{"c", "d", "a", "b"},
		},
	})

	sheet, layout := this.prepare(this.excel, this.sheetSuccess)
	result, err := JsonPack("1", this.lintOfData, sheet, layout)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(expected1), string(result))

	sheet, layout = this.prepare(this.excel, this.sheetSuccess)
	result, err = JsonPack("2", this.lintOfData, sheet, layout)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(expected2), string(result))

	sheet, layout = this.prepare(this.excel, this.sheetSuccess)
	result, err = JsonPack("3", this.lintOfData, sheet, layout)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), string(expected3), string(result))

	sheet, layout = this.prepare(this.excel, this.sheetFailed)
	_, err = JsonPack("1", this.lintOfData, sheet, layout)
	assert.NotNil(this.T(), err)

	sheet, layout = this.prepare(this.excel, this.sheetEmpty)
	_, err = JsonPack("1", this.lintOfData, sheet, layout)
	assert.NotNil(this.T(), err)
}

func (this *SuiteJsonPack) prepare(excelName, sheetName string) (sheet *excels.Sheet, layout *Layout) {
	excel := &excels.Excel{}
	assert.Nil(this.T(), excel.Open(excelName))
	sheet, err := excel.Get(sheetName)
	assert.Nil(this.T(), err)
	line, err := excel.GetLine(
		sheetName,
		this.lineOfTag,
		this.lineOfName,
		this.lineOfNote,
		this.lineOfField,
	)
	assert.Nil(this.T(), err)

	layout = NewLayout()
	lineTag := line[this.lineOfTag]
	lineName := line[this.lineOfName]
	lineNote := line[this.lineOfNote]
	lineField := line[this.lineOfField]
	assert.Nil(this.T(), layout.Set(lineTag, lineName, lineNote, lineField))
	return sheet, layout
}
