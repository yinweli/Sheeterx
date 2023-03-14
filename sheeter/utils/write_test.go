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
	testdata.TestData
	fileExist string
}

func (this *SuiteWrite) SetupSuite() {
	this.TBegin("test-utils-write", "write")
	this.fileExist = "exist.txt"
}

func (this *SuiteWrite) TearDownSuite() {
	this.TFinal()
}

func (this *SuiteWrite) TestFileName() {
	assert.Equal(this.T(), "test", FileName(filepath.Join("dir1", "dir2", "dir3", "test.txt")))
}

func (this *SuiteWrite) TestFileExist() {
	assert.True(this.T(), FileExist(this.fileExist))
	assert.False(this.T(), FileExist(this.Unknown))
}

func (this *SuiteWrite) TestWriteFile() {
	path := filepath.Join("path", "test.file")
	data := []byte("this is a string")

	assert.Nil(this.T(), WriteFile(path, data))
	this.CompareFile(this.T(), path, data)
}
