package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestLongArray(t *testing.T) {
	suite.Run(t, new(SuiteLongArray))
}

type SuiteLongArray struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteLongArray) SetupSuite() {
	this.TBegin("test-fields-longArray", "")
}

func (this *SuiteLongArray) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteLongArray) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"longArray", "[]long", "long[]"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeLongCs+sheeter.TypeArray, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeArray+sheeter.TypeLongGo, target.ToTypeGo())
}

func (this *SuiteLongArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}

func (this *SuiteLongArray) target() *LongArray {
	return &LongArray{}
}
