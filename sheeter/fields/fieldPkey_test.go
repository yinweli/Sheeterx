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
	testdata.TestEnv
}

func (this *SuitePkey) SetupSuite() {
	this.Change("test-fields-pkey")
}

func (this *SuitePkey) TearDownSuite() {
	this.Restore()
}

func (this *SuitePkey) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"pkey"}, target.Field())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.ToTypeGo())
}

func (this *SuitePkey) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "123456789", result)
}

func (this *SuitePkey) target() *Pkey {
	return &Pkey{}
}
