package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestFloat(t *testing.T) {
	suite.Run(t, new(SuiteFloat))
}

type SuiteFloat struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteFloat) SetupSuite() {
	this.TBegin("test-fields-float", "")
}

func (this *SuiteFloat) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteFloat) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"float"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeFloatCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeFloatGo, target.ToTypeGo())
}

func (this *SuiteFloat) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float32(0.123456), result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}

func (this *SuiteFloat) target() *Float {
	return &Float{}
}
