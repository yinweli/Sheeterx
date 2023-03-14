package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestString(t *testing.T) {
	suite.Run(t, new(SuiteString))
}

type SuiteString struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteString) SetupSuite() {
	this.TBegin("test-fields-string", "")
}

func (this *SuiteString) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteString) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"string"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeStringCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeStringGo, target.ToTypeGo())
}

func (this *SuiteString) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}

func (this *SuiteString) target() *String {
	return &String{}
}
