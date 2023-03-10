package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/sheeter/fields"
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
	this.Change("test-nameds-pkey")
}

func (this *SuitePkey) TearDownSuite() {
	this.Restore()
}

func (this *SuitePkey) TestPkey() {
	target := this.target()
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.PkeyCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.PkeyGo())
}

func (this *SuitePkey) target() *Pkey {
	return &Pkey{
		Pkey: &fields.Pkey{},
	}
}
