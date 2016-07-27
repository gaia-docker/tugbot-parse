package parse

import (
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJunitToTestSet(t *testing.T) {
	log.Info("TestJunitToTestSet")
	testSet, err := ToTestSet([]byte(getMockMochaXunit2()))
	assert.NoError(t, err)
	assert.Equal(t, FloatNumber(0.149), testSet.Time)
	assert.Equal(t, 5, len(testSet.Tests))
	assert.Equal(t, 5, testSet.Total, "Total")
	assert.Equal(t, 1, testSet.Failures, "Failures")
	assert.Equal(t, 1, testSet.Errors, "Errors")
	assert.Equal(t, 0, testSet.Skipped, "Skipped")
	failureTest := getFailureTest(testSet.Tests)
	assert.NotEqual(t, "", failureTest.Failure)
	assert.Equal(t, "Failed", failureTest.Status)
}

func getFailureTest(tests []Test) Test {
	var ret Test
	for _, currTest := range tests {
		if currTest.Failure != "" {
			ret = currTest
			break
		}
	}

	return ret
}
