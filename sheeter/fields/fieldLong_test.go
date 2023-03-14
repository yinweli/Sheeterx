package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestLong(t *testing.T) {
	suite.Run(t, new(SuiteLong))
}

type SuiteLong struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteLong) SetupSuite() {
	this.TBegin("test-fields-long", "")
}

func (this *SuiteLong) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteLong) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"long"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeLongCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeLongGo, target.ToTypeGo())
}

func (this *SuiteLong) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(this.Unknown)
	assert.NotNil(this.T(), err)
}

func (this *SuiteLong) target() *Long {
	return &Long{}
}
