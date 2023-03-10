package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestFloatArray(t *testing.T) {
	suite.Run(t, new(SuiteFloatArray))
}

type SuiteFloatArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteFloatArray) SetupSuite() {
	this.Change("test-fields-floatArray")
}

func (this *SuiteFloatArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteFloatArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"floatArray", "[]float", "float[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeFloatCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeFloatGo, target.ToTypeGo())
}

func (this *SuiteFloatArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float32{0.123, 0.456, 0.789}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteFloatArray) target() *FloatArray {
	return &FloatArray{}
}
