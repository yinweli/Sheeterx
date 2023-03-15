package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestJson(t *testing.T) {
	suite.Run(t, new(SuiteJson))
}

type SuiteJson struct {
	suite.Suite
	testdata.TestData
}

func (this *SuiteJson) SetupSuite() {
	this.TBegin("test-utils-json", "")
}

func (this *SuiteJson) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteJson) TestJsonMarshal() {
	datas := map[string]string{"data": "value"}
	bytes, _ := json.MarshalIndent(datas, sheeter.JsonPrefix, sheeter.JsonIdent)

	result, err := JsonMarshal(datas)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), bytes, result)
}
