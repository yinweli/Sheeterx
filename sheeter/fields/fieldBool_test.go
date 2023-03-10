package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestBool(t *testing.T) {
	suite.Run(t, new(SuiteBool))
}

type SuiteBool struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteBool) SetupSuite() {
	this.Change("test-fields-bool")
}

func (this *SuiteBool) TearDownSuite() {
	this.Restore()
}

func (this *SuiteBool) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"bool"}, target.Field())
	assert.Equal(this.T(), false, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeBoolCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeBoolGo, target.ToTypeGo())
}

func (this *SuiteBool) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, result)

	result, err = target.ToJsonValue("false")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, result)

	_, err = target.ToJsonValue("")
	assert.NotNil(this.T(), err)

	_, err = target.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteBool) target() *Bool {
	return &Bool{}
}
