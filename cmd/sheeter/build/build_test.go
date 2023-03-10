package build

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/sheeter"
	"github.com/yinweli/Sheeterx/sheeter/excels"
	"github.com/yinweli/Sheeterx/testdata"
)

func TestBuild(t *testing.T) {
	suite.Run(t, new(SuiteBuild))
}

type SuiteBuild struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteBuild) SetupSuite() {
	this.Change("test-cmd-build")
}

func (this *SuiteBuild) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteBuild) TestNewCommand() {
	assert.NotNil(this.T(), NewCommand())
}

func (this *SuiteBuild) TestExecute() {
	config := "config"
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Flags().Set(config, testdata.ConfigBuild))
	assert.Nil(this.T(), cmd.Execute())
	assert.FileExists(this.T(), filepath.Join(sheeter.CsPath, "RealData.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.CsPath, "Sheeter.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.GoPath, "sheeter.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, "realData.json"))
}
