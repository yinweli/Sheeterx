package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestCheck(t *testing.T) {
	suite.Run(t, new(SuiteCheck))
}

type SuiteCheck struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteCheck) SetupSuite() {
	this.Change("test-utils-check")
}

func (this *SuiteCheck) TearDownSuite() {
	this.Restore()
}

func (this *SuiteCheck) TestCheckKeyword() {
	assert.True(this.T(), CheckKeyword("value"))
	assert.False(this.T(), CheckKeyword("sheeter"))
	assert.False(this.T(), CheckKeyword("Sheeter"))
	assert.False(this.T(), CheckKeyword("SHEETER"))
}

func (this *SuiteCheck) TestCheckIgnore() {
	assert.True(this.T(), CheckIgnore(sheeter.TokenIgnore+"data"))
	assert.False(this.T(), CheckIgnore(testdata.UnknownStr))
}

func (this *SuiteCheck) TestCheckName() {
	assert.True(this.T(), CheckName("value"))
	assert.True(this.T(), CheckName("Value"))
	assert.True(this.T(), CheckName("value1"))
	assert.True(this.T(), CheckName("Value1"))
	assert.True(this.T(), CheckName("value_"))
	assert.True(this.T(), CheckName("_value"))
	assert.False(this.T(), CheckName(""))
	assert.False(this.T(), CheckName("0value"))
	assert.False(this.T(), CheckName("-value"))
	assert.False(this.T(), CheckName("value-"))
	assert.False(this.T(), CheckName("#value"))
	assert.False(this.T(), CheckName("value#"))
	assert.False(this.T(), CheckName("@value"))
	assert.False(this.T(), CheckName("value@"))
	assert.False(this.T(), CheckName("{value}"))
}

func (this *SuiteCheck) TestCheckTag() {
	assert.True(this.T(), CheckTag("a", "abc"))
	assert.True(this.T(), CheckTag("b", "abc"))
	assert.True(this.T(), CheckTag("c", "abc"))
	assert.True(this.T(), CheckTag("ab", "abc"))
	assert.True(this.T(), CheckTag("bc", "abc"))
	assert.True(this.T(), CheckTag("ac", "abc"))
	assert.False(this.T(), CheckTag("x", "abc"))
	assert.False(this.T(), CheckTag("i", sheeter.TokenIgnore))
}
