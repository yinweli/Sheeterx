package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestStringArray(t *testing.T) {
	suite.Run(t, new(SuiteStringArray))
}

type SuiteStringArray struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteStringArray) SetupSuite() {
	this.TBegin("test-fields-stringArray", "")
}

func (this *SuiteStringArray) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteStringArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"stringArray", "[]string", "string[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeStringCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeStringGo, target.ToTypeGo())
}

func (this *SuiteStringArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"ball", "book", "pack"}, result)
}

func (this *SuiteStringArray) target() *StringArray {
	return &StringArray{}
}
