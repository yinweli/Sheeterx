package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestBoolArray(t *testing.T) {
	suite.Run(t, new(SuiteBoolArray))
}

type SuiteBoolArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteBoolArray) SetupSuite() {
	this.Change("test-fields-boolArray")
}

func (this *SuiteBoolArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteBoolArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"boolArray", "[]bool", "bool[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeBoolCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeBoolGo, target.ToTypeGo())
}

func (this *SuiteBoolArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteBoolArray) target() *BoolArray {
	return &BoolArray{}
}
