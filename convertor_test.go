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
	failureTests := getFailureTests(testSet.Tests)
	assert.Equal(t, 1, len(failureTests))
	assert.Equal(t, "Failed", failureTests[0].Status)
}

func TestJunitToTestSet_SkippedTest(t *testing.T) {
	log.Info("TestJunitToTestSet_SkippedTest")
	testSet, err := ToTestSet([]byte(getMockMochaXunitWithSkippedTest()))
	assert.NoError(t, err)
	assert.Equal(t, FloatNumber(4.045), testSet.Time)
	assert.Equal(t, 5, len(testSet.Tests))
	assert.Equal(t, 5, testSet.Total, "Total")
	assert.Equal(t, 4, testSet.Failures, "Failures")
	assert.Equal(t, 4, testSet.Errors, "Errors")
	assert.Equal(t, 1, testSet.Skipped, "Skipped")
	assert.Equal(t, 1, len(getSkippedTests(testSet.Tests)))
}

func getFailureTests(tests []Test) []Test {
	var ret []Test
	for _, currTest := range tests {
		if currTest.Failure != "" {
			ret = append(ret, currTest)
		}
	}

	return ret
}

func getSkippedTests(tests []Test) []Test {
	var ret []Test
	for _, currTest := range tests {
		if currTest.Status == "Skipped" {
			ret = append(ret, currTest)
		}
	}

	return ret
}
