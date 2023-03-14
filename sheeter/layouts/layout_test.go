package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter/fields"
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

func (this *SuiteLayout) TestAdd() {
	target := NewLayout()
	assert.Nil(this.T(), target.Add("", "name1", "note1", &fields.Pkey{}))
	assert.Nil(this.T(), target.Add("", "name2", "note2", &fields.Int{}))
	assert.NotNil(this.T(), target.Add("", "name2", "note2", &fields.Int{}))
	assert.NotNil(this.T(), target.Add("", "name3", "note3", &fields.Pkey{}))
	assert.NotNil(this.T(), target.Add("", "name4", "note4", nil))
	assert.NotNil(this.T(), target.Add("", "name@", "note5", &fields.Int{}))
	assert.NotNil(this.T(), target.Add("", "", "note6", &fields.Int{}))
}

func (this *SuiteLayout) TestPack() {
	target := NewLayout()
	assert.Nil(this.T(), target.Add("123", "name1", "note1", &fields.Pkey{}))
	assert.Nil(this.T(), target.Add("12", "name2", "note2", &fields.Int{}))
	assert.Nil(this.T(), target.Add("12", "name3", "note3", &fields.IntArray{}))
	assert.Nil(this.T(), target.Add("13", "name4", "note4", &fields.String{}))
	assert.Nil(this.T(), target.Add("13", "name5", "note5", &fields.StringArray{}))

	data := []string{"1", "2", "1,2,3", "a", "a,b,c"}
	actual1 := map[string]interface{}{
		"name1": "1",
		"name2": int32(2),
		"name3": []int32{1, 2, 3},
		"name4": "a",
		"name5": []string{"a", "b", "c"},
	}
	actual2 := map[string]interface{}{
		"name1": "1",
		"name2": int32(2),
		"name3": []int32{1, 2, 3},
	}
	actual3 := map[string]interface{}{
		"name1": "1",
		"name4": "a",
		"name5": []string{"a", "b", "c"},
	}

	result, pkey, err := target.Pack("1", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "1", pkey)
	assert.Equal(this.T(), actual1, result)

	result, pkey, err = target.Pack("2", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "1", pkey)
	assert.Equal(this.T(), actual2, result)

	result, pkey, err = target.Pack("3", data)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "1", pkey)
	assert.Equal(this.T(), actual3, result)

	_, _, err = target.Pack("1", []string{"1", "@", "1,2,3", "a", "a,b,c"})
	assert.NotNil(this.T(), err)
}

func (this *SuiteLayout) TestLayout() {
	target := NewLayout()
	assert.Nil(this.T(), target.Add("123", "name1", "note1", &fields.Pkey{}))
	assert.Nil(this.T(), target.Add("12", "name2", "note2", &fields.Int{}))
	assert.Nil(this.T(), target.Add("12", "name3", "note3", &fields.IntArray{}))
	assert.Nil(this.T(), target.Add("13", "name4", "note4", &fields.String{}))
	assert.Nil(this.T(), target.Add("1", "name5", "note5", &fields.StringArray{}))

	assert.Len(this.T(), target.Layout("1"), 5)
	assert.Len(this.T(), target.Layout("2"), 3)
	assert.Len(this.T(), target.Layout("3"), 2)
	assert.Len(this.T(), target.Layout("4"), 0)
}

func (this *SuiteLayout) TestPkey() {
	target := NewLayout()
	assert.Nil(this.T(), target.Add("a", "name1", "note1", &fields.Pkey{}))
	assert.NotNil(this.T(), target.Pkey("a"))
	assert.Nil(this.T(), target.Pkey("b"))
}
