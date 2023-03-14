package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestInt(t *testing.T) {
	suite.Run(t, new(SuiteInt))
}

type SuiteInt struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteInt) SetupSuite() {
	this.TBegin("test-fields-int", "")
}

func (this *SuiteInt) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteInt) TestField() {
	target := &Int{}
	assert.Equal(this.T(), []string{"int"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeIntCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeIntGo, target.ToTypeGo())
}

func (this *SuiteInt) TestToJsonValue() {
	target := &Int{}

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(123456789), result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}
