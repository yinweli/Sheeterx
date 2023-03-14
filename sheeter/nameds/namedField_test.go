package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/sheeter/fields"
	"github.com/yinweli/Sheeterx/sheeter/layouts"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestField(t *testing.T) {
	suite.Run(t, new(SuiteField))
}

type SuiteField struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteField) SetupSuite() {
	this.TBegin("test-nameds-field", "")
}

func (this *SuiteField) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteField) TestName() {
	target := this.target()
	assert.Equal(this.T(), "Name", target.FieldName())
	assert.Equal(this.T(), "note", target.FieldNote())
	assert.Equal(this.T(), sheeter.TypePkeyCs, target.FieldTypeCs())
	assert.Equal(this.T(), sheeter.TypePkeyGo, target.FieldTypeGo())
}

func (this *SuiteField) target() *Field {
	return &Field{
		Data: &layouts.Data{
			Tag:   "",
			Name:  "name",
			Note:  "note",
			Field: &fields.Pkey{},
		},
	}
}
