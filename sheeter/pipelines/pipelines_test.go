package pipelines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeterx/testdata"
)

func TestPipeline(t *testing.T) {
	suite.Run(t, new(SuitePipeline))
}

type SuitePipeline struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuitePipeline) SetupSuite() {
	this.Change("test-pipelines-pipeline")
}

func (this *SuitePipeline) TearDownSuite() {
	this.Restore()
}

func (this *SuitePipeline) TestPipeline() {
	result, errs := Pipeline[int]("name", []int{0, 1}, []PipelineFunc[int]{
		func(material int, result chan any) error {
			result <- material
			return nil
		},
		func(material int, result chan any) error {
			return fmt.Errorf("err")
		},
	})
	assert.Len(this.T(), result, 2)
	assert.Len(this.T(), errs, 2)

	result, errs = Pipeline[int]("name", []int{}, []PipelineFunc[int]{})
	assert.Empty(this.T(), result)
	assert.Empty(this.T(), errs)
}
