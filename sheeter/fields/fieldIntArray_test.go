package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestIntArray(t *testing.T) {
	suite.Run(t, new(SuiteIntArray))
}

type SuiteIntArray struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteIntArray) SetupSuite() {
	this.Change("test-fields-intArray")
}

func (this *SuiteIntArray) TearDownSuite() {
	this.Restore()
}

func (this *SuiteIntArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"intArray", "[]int", "int[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeIntCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeIntGo, target.ToTypeGo())
}

func (this *SuiteIntArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int32{123, 456, 789}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteIntArray) target() *IntArray {
	return &IntArray{}
}
