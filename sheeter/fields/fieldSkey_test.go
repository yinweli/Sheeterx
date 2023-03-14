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
	testdata.TestData
}

func (this *SuiteSkey) SetupSuite() {
	this.TBegin("test-fields-skey", "")
}

func (this *SuiteSkey) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteSkey) TestField() {
	target := &Skey{}
	assert.Equal(this.T(), []string{"skey"}, target.Field())
	assert.Equal(this.T(), true, target.IsPkey())
	assert.Equal(this.T(), sheeter.TypeSkeyCs, target.ToTypeCs())
	assert.Equal(this.T(), sheeter.TypeSkeyGo, target.ToTypeGo())
}

func (this *SuiteSkey) TestToJsonValue() {
	target := &Skey{}

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}
