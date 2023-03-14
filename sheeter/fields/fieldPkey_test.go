package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestPkey(t *testing.T) {
	suite.Run(t, new(SuitePkey))
}

type SuitePkey struct {
	suite.Suite
	testdata.TestData
}

func (this *SuitePkey) SetupSuite() {
	this.TBegin("test-fields-pkey", "")
}

func (this *SuitePkey) TearDownSuite() {
	this.TFinal()
}

func (this *SuitePkey) TestField() {
	target := &Pkey{}
	assert.Equal(this.T(), []string{"pkey"}, target.Field())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.ToTypeGo())
}

func (this *SuitePkey) TestToJsonValue() {
	target := &Pkey{}

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "123456789", result)
}
