package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestLayout(t *testing.T) {
	suite.Run(t, new(SuiteLayout))
}

type SuiteLayout struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteLayout) SetupSuite() {
	this.TBegin("test-layouts-layout", "")
}

func (this *SuiteLayout) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteLayout) TestNewLayout() {
	assert.NotNil(this.T(), NewLayout())
}

func (this *SuiteLayout) TestSet() {
	target := NewLayout()
	assert.Nil(this.T(), target.Set(
		[]string{"", "123", "ignore", "12", "12", "13", "13"},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5"},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5"},
		[]string{"", "pkey", "", "int", "[]int", "string", "[]string"}))

	target = NewLayout()
	assert.NotNil(this.T(), target.Set(
		[]string{"", "123", "ignore", "12", "12", "13", "13"},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5"},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5"},
		[]string{"", "@", "", "int", "[]int", "string", "[]string"}))
}

func (this *SuiteLayout) TestPack() {
	target := NewLayout()
	assert.Nil(this.T(), target.Set(
		[]string{"", "123", "ignore", "12", "12", "13", "13"},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5"},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5"},
		[]string{"", "pkey", "", "int", "[]int", "string", "[]string"}))

	data := []string{"", "1", "", "2", "1,2,3", "a", "a,b,c"}
	dataIgnore := []string{sheeter.TokenIgnore, "1", "", "2", "1,2,3", "a", "a,b,c"}
	dataInvalid := []string{"", "1", "", "@", "1,2,3", "a", "a,b,c"}
	actual1 := map[string]interface{}{
		"name1": int32(1),
		"name2": int32(2),
		"name3": []int32{1, 2, 3},
		"name4": "a",
		"name5": []string{"a", "b", "c"},
	}
	actual2 := map[string]interface{}{
		"name1": int32(1),
		"name2": int32(2),
		"name3": []int32{1, 2, 3},
	}
	actual3 := map[string]interface{}{
		"name1": int32(1),
		"name4": "a",
		"name5": []string{"a", "b", "c"},
	}

	result, pkey, err := target.Pack("1", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(1), pkey)
	assert.Equal(this.T(), actual1, result)

	result, pkey, err = target.Pack("2", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(1), pkey)
	assert.Equal(this.T(), actual2, result)

	result, pkey, err = target.Pack("3", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(1), pkey)
	assert.Equal(this.T(), actual3, result)

	result, pkey, err = target.Pack("1", dataIgnore)
	assert.Nil(this.T(), err)
	assert.Nil(this.T(), pkey)
	assert.Nil(this.T(), result)

	_, _, err = target.Pack("1", dataInvalid)
	assert.NotNil(this.T(), err)
}

func (this *SuiteLayout) TestLayout() {
	target := NewLayout()
	assert.Nil(this.T(), target.Set(
		[]string{"", "123", "ignore", "12", "12", "13", "1"},
		[]string{"", "name1", "", "name2", "name3", "name4", "name5"},
		[]string{"", "note1", "", "note2", "note3", "note4", "note5"},
		[]string{"", "pkey", "", "int", "[]int", "string", "[]string"}))

	assert.Len(this.T(), target.Layout("1"), 5)
	assert.Len(this.T(), target.Layout("2"), 3)
	assert.Len(this.T(), target.Layout("3"), 2)
	assert.Len(this.T(), target.Layout("4"), 0)
}

func (this *SuiteLayout) TestPkey() {
	target := NewLayout()
	assert.Nil(this.T(), target.Set(
		[]string{"", "1"},
		[]string{"", "name1"},
		[]string{"", "note1"},
		[]string{"", "pkey"}))

	assert.NotNil(this.T(), target.Pkey("1"))
	assert.Nil(this.T(), target.Pkey("2"))
}

func (this *SuiteLayout) TestAdd() {
	target := NewLayout()
	assert.Nil(this.T(), target.add(0, "", "name1", "note1", "pkey"))
	assert.Nil(this.T(), target.add(0, "", "name2", "note2", "int"))
	assert.NotNil(this.T(), target.add(0, "", "name2", "note2", "int"))
	assert.NotNil(this.T(), target.add(0, "", "name3", "note3", "pkey"))
	assert.NotNil(this.T(), target.add(0, "", "name4", "note4", "@"))
	assert.NotNil(this.T(), target.add(0, "", "name@", "note5", "int"))
	assert.NotNil(this.T(), target.add(0, "", "", "note6", "int"))
}
