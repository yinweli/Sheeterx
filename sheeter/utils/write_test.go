package utils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/testdata"
)

func TestWrite(t *testing.T) {
	suite.Run(t, new(SuiteWrite))
}

type SuiteWrite struct {
	suite.Suite
	testdata.TestEnv
	dirReal string
}

func (this *SuiteWrite) SetupSuite() {
	this.Change("test-utils-write")
	this.dirReal = "write"
}

func (this *SuiteWrite) TearDownSuite() {
	this.Restore()
}

func (this *SuiteWrite) TestFileName() {
	assert.Equal(this.T(), "test", FileName(filepath.Join("dir1", "dir2", "dir3", "test.txt")))
}

func (this *SuiteWrite) TestFileExist() {
	assert.True(this.T(), FileExist(testdata.ConfigReal))
	assert.False(this.T(), FileExist(testdata.UnknownStr))
}

func (this *SuiteWrite) TestWriteFile() {
	path := filepath.Join("path", "test.file")
	data := []byte("this is a string")

	assert.Nil(this.T(), WriteFile(path, data))
	testdata.CompareFile(this.T(), path, data)
}
