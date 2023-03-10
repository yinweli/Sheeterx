package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestSkey(t *testing.T) {
	suite.Run(t, new(SuiteSkey))
}

type SuiteSkey struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteSkey) SetupSuite() {
	this.Change("test-fields-skey")
}

func (this *SuiteSkey) TearDownSuite() {
	this.Restore()
}

func (this *SuiteSkey) TestField() {
	target := this.target()
	assert.Equal(this.T(), []string{"skey"}, target.Field())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeSkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeSkeyGo, target.ToTypeGo())
}

func (this *SuiteSkey) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}

func (this *SuiteSkey) target() *Skey {
	return &Skey{}
}