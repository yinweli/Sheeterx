package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestDouble(t *testing.T) {
	suite.Run(t, new(SuiteDouble))
}

type SuiteDouble struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteDouble) SetupSuite() {
	this.TBegin("test-fields-double", "")
}

func (this *SuiteDouble) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteDouble) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"double"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeDoubleCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeDoubleGo, target.ToTypeGo())
}

func (this *SuiteDouble) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123456")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), 0.123456, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}

func (this *SuiteDouble) target() *Double {
	return &Double{}
}
